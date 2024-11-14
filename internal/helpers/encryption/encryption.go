package encryption

type Encryption interface {
	encrypt(password string) (string, error)
	decrypt(hashedPassword, password string) (bool, error)
}

func Encrypting(e Encryption, password string) (string, error) {
	return e.encrypt(password)
}

func Decrypting(e Encryption, hashedPassword string, password string) (bool, error) {
	return e.decrypt(hashedPassword, password)
}
