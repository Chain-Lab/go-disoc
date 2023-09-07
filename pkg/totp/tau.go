/**
  @author: decision
  @date: 2023/9/6
  @note:
**/

package totp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"strconv"
	"time"
)

// KeyGenTOTP TOTPGen() -> sk
// On input a security parameter 1^lambda, it generates a secret key $sk$.
func KeyGenTOTP() string {
	timestamp := time.Now().UnixMilli()

	sha := sha256.New()
	sha.Write([]byte(strconv.FormatInt(timestamp, 10)))
	digestBytes := sha.Sum(nil)
	base32Code := base32.StdEncoding.EncodeToString(digestBytes)
	return base32Code[:16]
}

// GenTOTP (sk, C_T) -> \tau
// On input the secret key sk and a time counter C_T, it computes the TOTP
// result \tau at counter C_T
func GenTOTP(key []byte, epoch []byte) []byte {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(epoch)
	hash := hmacSha1.Sum(nil)

	return hash
}

// VerifyTOTP (sk, \tau, C_T) -> {0, 1}
// On input the secret key sk, the time counter C_T and TOTP result \tau, it
// outputs true if and only if \tau = TOTPGen(sk, C_T).
func VerifyTOTP(key []byte, code []byte, epoch []byte) bool {
	correctCode := GenTOTP(key, epoch)

	return bytes.Equal(correctCode, code)
}
