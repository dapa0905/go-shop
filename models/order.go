package models

import (
	"go-shop/models/enum"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint
	TotalPrice float64
	Status     enum.OrderStatus
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}
