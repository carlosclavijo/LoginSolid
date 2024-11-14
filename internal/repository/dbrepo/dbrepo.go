package dbrepo

import (
	"database/sql"

	"github.com/carlosclavijo/loginsolid/internal/repository"
)

type dbRepo struct {
	DB *sql.DB
}

func NewDbRepo(conn *sql.DB) repository.Repository {
	return &dbRepo{
		DB: conn,
	}
}
