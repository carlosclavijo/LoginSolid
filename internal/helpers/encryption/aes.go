package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type Aes struct {
	Key string
}

func NewAes() Encryption {
	return &Aes{
		Key: "solidprincipleskeyvalidatorvalue",
	}
}

func (a *Aes) encrypt(password string) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		return "", err
	}
	plainText := []byte(password)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.RawStdEncoding.EncodeToString(cipherText), nil
}

func (a *Aes) decrypt(hashedPassword, password string) (bool, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, err
	}
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		return false, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext) == password, nil
}
