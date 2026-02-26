package domain

import (
	"time"
)

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Description string
	Price float64
	Image string
	CreatedAt time.Time
	UpdateAt time.Time
}