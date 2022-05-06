package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func EncryptOAEP(text string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	rsaPublicKey := pubInterface.(*rsa.PublicKey)
	secretMessage := []byte(text)
	rng := rand.Reader
	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, rsaPublicKey, secretMessage, nil)
	if err != nil {
		return "", nil
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	return ciphertext, nil
}

// DecryptOAEP 解密
func DecryptOAEP(ciphertext string, privateKey []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	rsaPrivateKey := privateInterface

	cipherData, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, rsaPrivateKey, cipherData, nil)
	if err != nil {
		return "", nil
	}

	return string(plaintext), nil
}
