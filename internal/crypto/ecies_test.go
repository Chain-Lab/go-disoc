/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	hexPrv := "3a07592b98bbde88c7be46d8e818b75d506ea08b3e52268e1e222775d0c3e2cc"
	hexPub := "03c45b85819cec7e05f606a85bb86643c188861b22c40eaaf91800d1428be100b1"

	message := []byte("message")

	ciphertext, err := EncryptData(message, hexPub)
	if err != nil {
		t.Errorf("Encrypt message failed")
	}

	plaintext, err := DecryptData(ciphertext, hexPrv)
	if err != nil {
		t.Fatal("Decrypt ciphertext failed")
		return
	}

	if !bytes.Equal(plaintext, message) {
		t.Fatal("Message and plaintext not equal")
		return
	}
}
