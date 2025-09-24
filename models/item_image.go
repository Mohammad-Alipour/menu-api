package models

import (
	"time"
)

type ItemImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemID    uint      `gorm:"not null" json:"item_id"`
	URL       string    `gorm:"type:varchar(255);not null" json:"url"`
	IsPrimary bool      `gorm:"default:false" json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`

	Item Item `gorm:"foreignKey:ItemID" json:"item,omitempty"`
}
