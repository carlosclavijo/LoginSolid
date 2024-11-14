package handlers

import (
	"github.com/carlosclavijo/loginsolid/internal/database"
	"github.com/carlosclavijo/loginsolid/internal/repository"
	"github.com/carlosclavijo/loginsolid/internal/repository/dbrepo"
)

type Repository struct {
	Db repository.Repository
}

var Repo *Repository

func NewRepo(db *database.DB) *Repository {
	return &Repository{
		Db: dbrepo.NewDbRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
