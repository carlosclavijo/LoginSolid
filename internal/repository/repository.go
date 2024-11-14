package repository

import "github.com/carlosclavijo/loginsolid/internal/models"

type Repository interface {
	CreateUser(r models.User) (models.User, error)
	GetUser(username string) (models.User, error)
	InsertLog(level, message, file string) error
}
