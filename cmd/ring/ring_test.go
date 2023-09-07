/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/athanorlabs/go-dleq/types"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/noot/ring-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"testing"
)

func TestRing(t *testing.T) {
	loadConfig()
	sha2 := sha256.New()
	sha2.Write([]byte("test"))
	digest := sha2.Sum(nil)

	sign, _ := RingSignature(digest)

	ringSign := ring.RingSig{}
	_ = ringSign.Deserialize(ring.Secp256k1(), sign)
	verified := checkPublicKeys(ringSign.PublicKeys())
	if !verified {
		t.Fatal("Group check failed.")
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
	groupKey = groupKey[:len(groupKey)-1]
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

func checkPublicKeys(keys []types.Point) bool {
	publicKeys := config.Strings("group")
	for idx := range keys {
		bytesPublicKey := keys[idx].Encode()
		hexPublicKey := hex.EncodeToString(bytesPublicKey)

		if !slices.Contains(publicKeys, hexPublicKey) {
			fmt.Printf("%d-th %s not exists\n", idx, hexPublicKey)
			return false
		}
	}

	return true
}
