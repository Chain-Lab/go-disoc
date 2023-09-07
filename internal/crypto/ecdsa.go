/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignatureMessage(data []byte, prvHex string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(prvHex)

	if err != nil {
		return nil, err
	}

	sha2 := sha256.New()
	sha2.Write(data)
	digest := sha2.Sum(nil)

	sign, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return nil, err
	}

	return sign[:len(sign)-1], nil
}

func VerifySignature(data []byte, pubHex string, signature []byte) (bool, error) {
	bytePublicKey, err := hex.DecodeString(pubHex)

	if err != nil {
		return false, err
	}

	publicECDSAKey, err := crypto.DecompressPubkey(bytePublicKey)

	if err != nil {
		return false, err
	}

	fullBytesPublic := crypto.FromECDSAPub(publicECDSAKey)

	sha2 := sha256.New()
	sha2.Write(data)
	digest := sha2.Sum(nil)

	return crypto.VerifySignature(fullBytesPublic, digest, signature), nil
}
