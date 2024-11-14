package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/loginsolid/internal/helpers/encryption"
	"github.com/carlosclavijo/loginsolid/internal/helpers/logger"
	"github.com/carlosclavijo/loginsolid/internal/models"
)

var encryptingState int = 0
var loggingState int = 0

func (m *Repository) PostUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	l := LoggingMethod(m)
	if err != nil {
		l.ErrorLog("Marshalling Json", err)
		return
	}
	crypting := CryptingMethod()
	User.Password, err = encryption.Encrypting(crypting, User.Password)
	if err != nil {
		l.ErrorLog("Encrypting password", err)
		return
	}
	User, err = m.Db.CreateUser(User)
	if err != nil {
		l.ErrorLog("Creating User", err)
		return
	}
	l.InfoLog("User Register", "User", User)
}

func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	l := LoggingMethod(m)
	if err != nil {
		l.ErrorLog("Marshalling Json", err)
		return
	}
	crypting := CryptingMethod()
	var u models.User
	u, err = m.Db.GetUser(User.Username)
	if err != nil {
		l.ErrorLog("Getting password", err)
		return
	}
	isLog, err := encryption.Decrypting(crypting, u.Password, User.Password)
	if err != nil {
		l.ErrorLog("Decrypting password", err)
		return
	}
	if !isLog {
		l.WarningLog("Incorrect password")
		return
	}
	l.InfoLog("Login Success", "User", u)
}

func (m *Repository) ChangeEncryption(w http.ResponseWriter, r *http.Request) {
	l := LoggingMethod(m)
	var method string
	if encryptingState < 2 {
		encryptingState++
	} else {
		encryptingState = 0
	}
	switch encryptingState {
	case 1:
		method = "AES"
	case 2:
		method = "BCrypts"
	case 3:
		method = "RSA"
	}
	l.InfoLog("Change Encryption", "Method", method)
}

func (m *Repository) ChangeLogging(w http.ResponseWriter, r *http.Request) {
	l := LoggingMethod(m)
	var method string
	if loggingState < 2 {
		loggingState++
	} else {
		loggingState = 0
	}
	switch loggingState {
	case 1:
		method = "Json"
	case 2:
		method = "DB"
	case 3:
		method = "TCP"
	}
	l.InfoLog("Change Logging", "Method", method)
}

func CryptingMethod() encryption.Encryption {
	switch encryptingState {
	case 0:
		return encryption.NewAes()
	case 1:
		return encryption.NewBcrypt()
	case 2:
		return encryption.NewRsa()
	}
	return nil
}

func LoggingMethod(m *Repository) logger.Logger {
	switch encryptingState {
	case 0:
		return logger.NewJsonLog()
	case 1:
		return logger.NewDbLog(m.Db)
		/*case 2:
		return encryption.NewRsa()*/
	}
	return nil
}
