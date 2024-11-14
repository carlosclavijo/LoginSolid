package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/loginsolid/internal/models"
)

func (m *dbRepo) CreateUser(r models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO users(username, password) VALUES($1, $2) RETURNING *`
	var u models.User
	err := m.DB.QueryRowContext(ctx, stmt, r.Username, r.Password).Scan(&u.UserId, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *dbRepo) GetUser(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT * FROM users WHERE username = $1`
	var u models.User
	err := m.DB.QueryRowContext(ctx, stmt, username).Scan(&u.UserId, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *dbRepo) InsertLog(level, message, file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO logs(level, message, file) VALUES($1, $2, $3) RETURNING *`
	_, err := m.DB.ExecContext(ctx, stmt, level, message, file)
	return err
}
