/**
  @author: decision
  @date: 2023/9/7
  @note: Simulate the node in DISOC
**/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/athanorlabs/go-dleq/types"
	crypto "github.com/chain-lab/go-disoc/internal/crypto"
	"github.com/chain-lab/go-disoc/internal/ethereum"
	"github.com/chain-lab/go-disoc/pkg/totp"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/noot/ring-go"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	loadConfig()

	// Generate AES temporary secret key
	key, err := crypto.GenerateRandomKey()
	log.WithField("timestamp", time.Now().UnixMilli()).
		Infoln("Node start require data on DON")

	log.Infof("Generate temporary key: %s", hex.EncodeToString(key))

	// Read totp secert key and interval from config
	totpKey := config.String("source.totp")
	interval := config.Int("source.interval")
	if err != nil {
		log.WithError(err).Errorln("Generate random key failed.")
		return
	}

	// Encrypt AES key with public key of data source
	encryptedKey, err := crypto.EncryptData(key, config.String("source.public"))
	if err != nil {
		log.WithError(err).Errorln("Encrypt temporary key failed.")
		return
	}

	epochSeconds := time.Now().Unix()
	epoch := totp.Int64ToBytes(epochSeconds / int64(interval))
	tau, err := totp.GenTOTP(totpKey, epoch)

	// Calculate digest of (encryptedKey, tau) by SHA-256
	sha2 := sha256.New()
	sha2.Write(encryptedKey)
	sha2.Write(tau)
	digest := sha2.Sum(nil)

	// Calculate ring signature
	rsign, err := RingSignature(digest)

	if err != nil {
		log.WithError(err).Errorln("Ring signature failed.")
	}

	// Create ethereum client with provider and contract address
	client, err := ethereum.NewEthereumClient(
		config.String("ethereum.provider"),
		config.String("ethereum.contract"))
	if err != nil {
		log.WithError(err).Errorln("Start ethereum client failed.")
		return
	}

	// Send node request, the process of subsection 4.2
	txHash, err := client.CallContractRequireData(encryptedKey, tau, rsign)
	log.Infof("Transaction send time stamp: %d", time.Now().UnixMilli())
	if err != nil {
		log.WithError(err).Errorln("Send transaction failed.")
		return
	}
	log.WithFields(log.Fields{
		"transaction": *txHash,
		"timestamp":   time.Now().UnixMilli(),
	}).Infoln("Transaction send.")

	// Monitor contract address with encryptedKey
	// the process of subsection 4.4
	event, err := client.ListenContractEvent(encryptedKey, nil)
	log.Infof("Receive event time stamp: %d", time.Now().UnixMilli())
	if err != nil {
		log.WithError(err).Errorln("Listen event failed.")
		return
	}

	log.Infof("Receive event: (%s, %s, %s)\n",
		hex.EncodeToString(event.EncryptedKey),
		hex.EncodeToString(event.EncryptedData),
		hex.EncodeToString(event.Signature))

	// Decrypt data from event.EncryptedData
	data, err := crypto.AESDecrypt(key, event.EncryptedData)
	if err != nil {
		log.WithError(err).Errorln("AES Decrypt failed.")
		return
	}

	log.WithField("timestamp", time.Now().UnixMilli()).
		Infof("Receive data: %s\n", hex.EncodeToString(data))

	verified, err := crypto.VerifySignature(data,
		config.String("source.public"), event.Signature)
	if err != nil || !verified {
		log.Warningln("Signature verify failed.")
	}
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)

	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		log.WithError(err).Errorln("Load config file failed.")
	}
}

func RingSignature(digest []byte) ([]byte, error) {
	curve := ring.Secp256k1()

	hexPrv := config.String("node.prv")
	//hexPub := config.String("node.pub")

	groupKey := config.Strings("group")
	publicKeys := make([]types.Point, len(groupKey))

	bytesPrv, err := hex.DecodeString(hexPrv)
	if err != nil {
		log.WithError(err).Errorln("Decode bytes from hex private key failed.")
		return nil, err
	}
	scalar := curve.ScalarFromBytes([32]byte(bytesPrv))

	for idx := range groupKey {
		bytesPub, err := hex.DecodeString(groupKey[idx])

		if err != nil {
			log.WithError(err).Fatal("Decode key from hex failed.")
			return nil, err
		}

		point, err := curve.DecodeToPoint(bytesPub)
		if err != nil {
			log.WithError(err).Errorln("Decode public key to point failed.")
			return nil, err
		}

		publicKeys[idx] = point
	}

	ringKey, err := ring.NewKeyRingFromPublicKeys(
		ring.Secp256k1(),
		publicKeys,
		scalar,
		len(publicKeys))
	if err != nil {
		log.WithError(err).Errorln("Create new key ring failed.")
		return nil, err
	}

	sign, err := ringKey.Sign([32]byte(digest), scalar)
	if err != nil {
		log.WithError(err).Errorln("Ring signature failed.")
		return nil, err
	}

	serializedSign, err := sign.Serialize()
	if err != nil {
		log.WithError(err).Errorln("Signature serialize failed.")
		return nil, err
	}

	return serializedSign, nil
}
