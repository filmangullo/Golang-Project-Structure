package models

import "time"

type Article struct {
	ID          uint64  `gorm:"primaryKey"`
	Title       string  `gorm:"type:varchar(255);not null"`
	Slug        string  `gorm:"type:varchar(255);unique;not null"`
	Content     string  `gorm:"type:longtext;not null"`
	Author      *string `gorm:"type:varchar(255);default:null"`
	Category    *string `gorm:"type:varchar(255);default:null"`
	Tags        *string `gorm:"size:155"`
	IsPublished *bool   `gorm:"not null;default:true"`
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
