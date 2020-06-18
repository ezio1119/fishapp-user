package domain

import "time"

type User struct {
	ID                int64 `json:"id"`
	Email             string
	Password          string `gorm:"-"`
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
