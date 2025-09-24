package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:varchar(20);default:'user'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
