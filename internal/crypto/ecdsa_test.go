/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package crypto

import (
	"testing"
)

func TestSignatureMessage(t *testing.T) {
	data := []byte("message")
	hexPrv := "3a07592b98bbde88c7be46d8e818b75d506ea08b3e52268e1e222775d0c3e2cc"
	hexPub := "03c45b85819cec7e05f606a85bb86643c188861b22c40eaaf91800d1428be100b1"
	signature, _ := SignatureMessage(data, hexPrv)

	verified, _ := VerifySignature(data, hexPub, signature)
	if !verified {
		t.Fatal("Signature verify failed.")
	}
}
