package model

import "time"

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price int
	Stock int

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
