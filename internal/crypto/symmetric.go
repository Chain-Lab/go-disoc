/**
  @author: decision
  @date: 2023/9/7
  @note:
**/

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
)

func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func GenerateRandomKey() ([]byte, error) {
	buf := make([]byte, 32)

	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func AESEncrypt(key, data []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	data = PKCS5Padding(data, cipher.BlockSize())

	if err != nil {
		return nil, err
	}

	bytesCipher := make([]byte, len(data))
	cipher.Encrypt(bytesCipher, data)

	return bytesCipher, nil
}

func AESDecrypt(key, ciphertext []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	bytesPlain := make([]byte, len(ciphertext))
	cipher.Decrypt(bytesPlain, ciphertext)

	bytesPlain = PKCS5UnPadding(bytesPlain)
	return bytesPlain, nil
}
