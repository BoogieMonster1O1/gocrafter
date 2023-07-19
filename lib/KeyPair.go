package lib

import (
	"crypto/ed25519"
	"crypto/rand"
	"os"
)

var (
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
)

func GetPublicKey() (ed25519.PublicKey, error) {
	if publicKey != nil {
		return publicKey, nil
	}

	publicKey, _, err := readKeyPair()
	if err != nil {
		return nil, err
	}

	return ed25519.PublicKey(publicKey), nil
}

func GetPrivateKey() (ed25519.PrivateKey, error) {
	if privateKey != nil {
		return privateKey, nil
	}

	_, privateKey, err := readKeyPair()
	if err != nil {
		return nil, err
	}

	return ed25519.PrivateKey(privateKey), nil
}

func GenerateKeyPair() error {
	pKey, pubKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	if err := writeKeyPair(ed25519.PrivateKey(pKey), ed25519.PublicKey(pubKey)); err != nil {
		return err
	}

	publicKey = ed25519.PublicKey(pubKey)
	privateKey = ed25519.PrivateKey(pKey)

	return nil
}

func readKeyPair() (ed25519.PrivateKey, ed25519.PublicKey, error) {
	privateKeyPath := "./run/id_rsa"
	publicKeyPath := "./run/id_rsa.pub"

	// Check if the key files exist
	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		// Generate new key pair if files don't exist
		if err := GenerateKeyPair(); err != nil {
			return nil, nil, err
		}
	}

	// Read the existing key pair files
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, publicKey, nil
}

func writeKeyPair(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) error {
	privateKeyPath := "./run/id_rsa"
	publicKeyPath := "./run/id_rsa.pub"

	if err := os.WriteFile(privateKeyPath, privateKey, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(publicKeyPath, publicKey, 0644); err != nil {
		return err
	}

	return nil
}
