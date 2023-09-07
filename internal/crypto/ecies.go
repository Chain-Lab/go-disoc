/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package crypto

import (
	ecies "github.com/ecies/go/v2"
)

func EncryptData(data []byte, pubHex string) ([]byte, error) {
	publicKey, err := ecies.NewPublicKeyFromHex(pubHex)

	if err != nil {
		return nil, err
	}

	ciphertext, err := ecies.Encrypt(publicKey, data)

	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func DecryptData(ciphertext []byte, prvHex string) ([]byte, error) {
	//privateECDSAKey, err := crypto.HexToECDSA(prvHex)
	//if err != nil {
	//	return nil, err
	//}

	privateKey, err := ecies.NewPrivateKeyFromHex(prvHex)
	if err != nil {
		return nil, err
	}

	plainText, err := ecies.Decrypt(privateKey, ciphertext)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
