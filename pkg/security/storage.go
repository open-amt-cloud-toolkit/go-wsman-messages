package security

import (
	"github.com/99designs/keyring"
)

// Storage struct to hold the service name and keyring interface

// NewStorage function to create a new Storage instance with a keyring interface (for testing).
func NewStorage(serviceName string, kr Keyring) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     kr,
	}
}

func NewKeyRingStorage(serviceName string) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     RealKeyring{},
	}
}

// SetKeyValue method to set a key-value pair in the keyring.
func (s Storage) SetKeyValue(key, value string) error {
	err := s.Keyring.Set(s.ServiceName, keyring.Item{
		Key:  key,
		Data: []byte(value),
	})

	return err
}

// GetKeyValue method to get a value from the keyring by key.
func (s Storage) GetKeyValue(key string) (string, error) {
	data, err := s.Keyring.Get(s.ServiceName, key)

	return string(data.Data), err
}

// Set method to set a key-value pair in the real keyring.
func (r RealKeyring) Set(serviceName string, item keyring.Item) error {
	kr, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})
	if err != nil {
		return err
	}

	return kr.Set(item)
}

// Get method to get a value from the real keyring by key.
func (r RealKeyring) Get(serviceName, key string) (keyring.Item, error) {
	kr, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})
	if err != nil {
		return keyring.Item{}, err
	}

	return kr.Get(key)
}
