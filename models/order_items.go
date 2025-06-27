package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint
	Price     float64
}
