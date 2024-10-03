package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt encrypts a string.
func (c Crypto) Encrypt(plainText []byte, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (c Crypto) GenerateKey() string {
	key := make([]byte, 24) // 24 bytes will result in a 32-character base64 string

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(key)
}
