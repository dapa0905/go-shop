package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	User      User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}
