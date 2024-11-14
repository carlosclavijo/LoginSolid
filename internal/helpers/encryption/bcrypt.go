package encryption

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type BCrypt struct{}

func NewBcrypt() Encryption {
	return &BCrypt{}
}

func (b *BCrypt) encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b *BCrypt) decrypt(hashedPassword, password string) (bool, error) {
	log.Println("HP", hashedPassword)
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
