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
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Int64ToBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

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
func GenTOTP(key string, epoch []byte) ([]byte, error) {
	secret, err := base32.StdEncoding.DecodeString(key)

	if err != nil {
		log.WithError(err).Errorln("Decode secret key from hex failed.")
		return nil, err
	}

	hmacSha1 := hmac.New(sha1.New, secret)
	hmacSha1.Write(epoch)
	hash := hmacSha1.Sum(nil)

	return hash, nil
}

// VerifyTOTP (sk, \tau, C_T) -> {0, 1}
// On input the secret key sk, the time counter C_T and TOTP result \tau, it
// outputs true if and only if \tau = TOTPGen(sk, C_T).
func VerifyTOTP(key string, code []byte, epoch []byte) bool {
	correctCode, err := GenTOTP(key, epoch)

	if err != nil {
		return false
	}

	return bytes.Equal(correctCode, code)
}
