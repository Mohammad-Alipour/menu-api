package models

import (
	"time"
)

type Item struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int       `gorm:"not null" json:"price"`
	Active      bool      `gorm:"default:true" json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Category   Category    `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ItemImages []ItemImage `gorm:"foreignKey:ItemID" json:"images,omitempty"`
}
