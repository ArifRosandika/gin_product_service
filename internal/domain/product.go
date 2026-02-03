package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID int
	Name string
	Description string
	Price float64
	Image string
}