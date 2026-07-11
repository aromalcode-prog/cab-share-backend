package models

import "time"

type User struct {
	ID           uint
	Name         string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
