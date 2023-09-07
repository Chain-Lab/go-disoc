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

func TestAESEncryptAndDecrypt(t *testing.T) {
	message := []byte("message")
	key, err := GenerateRandomKey()
	
	if err != nil {
		t.Fatal("Generate key failed.")
		return
	}

	cipher, err := AESEncrypt(key, message)
	if err != nil {
		t.Fatal("Encrypt data failed.")
		return
	}

	plaintext, err := AESDecrypt(key, cipher)
	if err != nil {
		t.Fatal("Decrypt cipher to data failed.")
		return
	}

	if !bytes.Equal(plaintext, message) {
		t.Fatal("Message before and after encrypt not equal.")
		return
	}
}
