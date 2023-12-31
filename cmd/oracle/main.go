/**
  @author: decision
  @date: 2023/9/7
  @note: Simulate the oracle and DON in DISOC
**/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/athanorlabs/go-dleq/types"
	"github.com/chain-lab/go-disoc/internal/crypto"
	"github.com/chain-lab/go-disoc/internal/ethereum"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/noot/ring-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
)

const (
	eventChanSize = 128
)

func main() {
	loadConfig()

	// Read totp secert key and interval from config
	interval := config.Int64("source.interval", 300)
	secret := config.String("source.totp")
	eventChan := make(chan *ethereum.EthereumDataRequest, eventChanSize)

	hexPrv := config.String("source.private")

	// Create ethereum client with provider and contract address
	client, err := ethereum.NewEthereumClient(
		config.String("ethereum.provider"),
		config.String("ethereum.contract"))
	if err != nil {
		log.WithError(err).Errorln("Start ethereum client failed.")
		return
	}

	// Start monitor contract events
	go client.OracleListenContractEvent(interval, secret, eventChan)
	log.Infoln("Start monitor events...")

	for {
		select {
		// Get event from event channel
		case event := <-eventChan:
			log.WithField("timestamp", time.Now().UnixMilli()).Infoln(
				"Contract request event received.")
			key, err := crypto.DecryptData(event.EncryptedKey, hexPrv)
			if err != nil {
				log.WithError(err).Errorln("Decrypto key failed.")
			}

			// Calculate digest in (encryptedKey, tau) with SHA-256
			// The process in subsection 4.2
			sha2 := sha256.New()
			sha2.Write(event.EncryptedKey)
			sha2.Write(event.Tau)
			digest := sha2.Sum(nil)

			// Verify the ring is correct
			ringSign := ring.RingSig{}
			ringSign.Deserialize(ring.Secp256k1(), event.RSignature)
			if !checkPublicKeys(ringSign.PublicKeys()) || !ringSign.Verify(
				[32]byte(digest)) {
				log.Warningln("Check group failed.")
				continue
			}

			// Simulate obtain data from DON
			data := strconv.Itoa(time.Now().Nanosecond())
			log.Infof("Waiting for message %s send.", data)
			signature, err := crypto.SignatureMessage([]byte(data), hexPrv)
			if err != nil {
				log.Warningln("Sign data failed.")
				continue
			}

			// Obtain AES key from encryptedKey by decrypt by private key
			encryptedData, err := crypto.AESEncrypt(key, []byte(data))

			if err != nil {
				log.Warningln("Encrypt data failed.")
				continue
			}

			// Send transaction, the data on-chaining process in subsection 4.4
			txHash, err := client.CallContractResponseData(event.EncryptedKey,
				encryptedData, signature)
			if err != nil {
				log.Warningln("Send contract response failed.")
				continue
			}

			log.WithFields(log.Fields{
				"transaction": *txHash,
				"timestamp":   time.Now().UnixMilli(),
			}).Infoln("Respond transaction send.")
		}
	}
}

func checkPublicKeys(keys []types.Point) bool {
	publicKeys := config.Strings("group")
	for idx := range keys {
		bytesPublicKey := keys[idx].Encode()
		hexPublicKey := hex.EncodeToString(bytesPublicKey)

		if !slices.Contains(publicKeys, hexPublicKey) {
			log.Warningf("%d th public key %s not exists", idx, hexPublicKey)
			return false
		}
	}

	return true
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)

	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		log.WithError(err).Errorln("Load config file failed.")
	}
}
