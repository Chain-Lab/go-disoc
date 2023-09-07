/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestCryptoCompress(t *testing.T) {
	hexPublic := "03c45b85819cec7e05f606a85bb86643c188861b22c40eaaf91800d1428be100b1"

	bytesPub, err := hex.DecodeString(hexPublic)

	if err != nil {
		t.Errorf("Hex decode failed.")
		return
	}

	pub, err := crypto.DecompressPubkey(bytesPub)

	if err != nil {
		t.Fatal("Decompress public key failed.")
		return
	}

	prv, err := crypto.HexToECDSA(
		"3a07592b98bbde88c7be46d8e818b75d506ea08b3e52268e1e222775d0c3e2cc")

	if err != nil {
		t.Errorf("Hex private key to ecdsa failed.")
		return
	}

	if !prv.PublicKey.Equal(pub) {
		t.Fatal("Public key not equal.")
	}
}

func TestGenerateKeyPair(t *testing.T) {
	GenerateKeyPair()
}
