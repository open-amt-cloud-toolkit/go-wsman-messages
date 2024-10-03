package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/config"
)

// Decrypt ciphertext using AES-GCM with the provided key.
func (c Crypto) Decrypt(cipherText string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, errors.New("cipher text too short")
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// Read encrypted data from file and decrypt it.
func (c Crypto) ReadAndDecryptFile(filePath string, key []byte) (config.Configuration, error) {
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return config.Configuration{}, err
	}

	decryptedData, err := c.Decrypt(string(encryptedData), key)
	if err != nil {
		return config.Configuration{}, err
	}

	var configuration config.Configuration

	err = yaml.Unmarshal(decryptedData, &configuration)
	if err != nil {
		return config.Configuration{}, err
	}

	return configuration, nil
}
