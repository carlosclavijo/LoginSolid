package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UserId    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
