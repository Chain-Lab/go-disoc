/**
  @author: decision
  @date: 2023/9/6
  @note:
**/

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/chain-lab/go-disoc/pkg/totp"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateTotpKey() {
	key := totp.KeyGenTOTP()
	fmt.Printf("Genrate TOTP key is: %s\n", key)
}

func GenerateKeyPair() {
	//curve := elliptic.P256()
	prv, err := crypto.GenerateKey()

	if err != nil {
		fmt.Printf("Generate key pair failed with error %s", err)
		return
	}

	prvHex := hex.EncodeToString(prv.D.Bytes())
	bytePub := crypto.CompressPubkey(&prv.PublicKey)
	pubHex := hex.EncodeToString(bytePub)

	fmt.Printf("Generate keypair result:\n")
	fmt.Printf("Private Key: %s\n", prvHex)
	fmt.Printf("Public Key: %s\n", pubHex)
}

func GenerateGroupKeys(size int) {
	fmt.Printf("Generate group key result:\n")
	for idx := 0; idx < size; idx++ {
		prv, err := crypto.GenerateKey()

		if err != nil {
			fmt.Printf("Generate key pair failed with error %s", err)
			return
		}

		bytePub := crypto.CompressPubkey(&prv.PublicKey)
		pubHex := hex.EncodeToString(bytePub)
		fmt.Printf("- \"%s\"\n", pubHex)
	}

}
