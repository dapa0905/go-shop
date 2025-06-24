package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"default:1"`
}
