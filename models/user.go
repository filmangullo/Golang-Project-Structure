package models

import "time"

type User struct {
	ID              uint64     `db:"id" gorm:"primaryKey"`
	Name            string     `db:"name" gorm:"type:varchar(255);not null"`
	Email           string     `db:"email" gorm:"type:varchar(255);unique;not null"`
	Password        string     `db:"password" gorm:"type:varchar(255);not null"`
	Role            *string    `db:"role" gorm:"type:varchar(55);default:null"`
	EmailVerifiedAt *time.Time `db:"email_verified_at" gorm:"type:timestamp;default:null"`
	RememberToken   *string    `db:"remember_token" gorm:"size:155"`
	IsActive        *bool      `db:"is_active" gorm:"not null;default:true"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
}
