package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"os"
)

type Rsa struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

func NewRsa() Encryption {
	pemFile, err := os.ReadFile("pkg/keys/private_key.pem")
	if err != nil {
		log.Panic(err)
	}
	block, _ := pem.Decode(pemFile)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Panic(err)
	}
	return &Rsa{
		PrivateKey: key,
		PublicKey:  key.PublicKey,
	}
}

func (r *Rsa) encrypt(password string) (string, error) {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &r.PublicKey, []byte(password), label)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (r *Rsa) decrypt(hashedPassword, password string) (bool, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, err
	}
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, r.PrivateKey, ciphertext, label)
	if err != nil {
		return false, err
	}
	return string(plaintext) == password, nil
}
