package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null;index:idx_cart_product,unique"`
	ProductID uint    `gorm:"not null;index:idx_cart_product,unique"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"default:1"`
}
