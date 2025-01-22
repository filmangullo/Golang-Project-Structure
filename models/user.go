package models

import "time"

type User struct {
	ID        uint   `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Password  string
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
