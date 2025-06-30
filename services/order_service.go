package services

import (
	"go-shop/config"
	"go-shop/dtos"
	"go-shop/models"
	"go-shop/models/enum"

	"gorm.io/gorm"
)

func CreateOrder(userID uint, req dtos.CreateOrderRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {

		// 장바구니 조회
		var cart models.Cart
		if err := tx.Where("user_id = ?", userID).First(&cart).Error; err != nil {
			return err
		}

		// 장바구니 아이템 -> OrderItem 으로 변환
		var totalPrice float64
		var orderItems []models.OrderItem
		for _, cartItem := range cart.CartItems {
			product, err := GetProductByID(cartItem.ProductID)
			if err != nil {
				return err
			}

			subtotal := float64(cartItem.Quantity) * product.Price
			totalPrice += subtotal

			orderItem := models.OrderItem{
				ProductID: cartItem.ProductID,
				Quantity:  cartItem.Quantity,
				Price:     product.Price,
			}

			orderItems = append(orderItems, orderItem)

		}

		// Order 생성 및 저장
		order := models.Order{
			UserID:     userID,
			TotalPrice: totalPrice,
			Status:     enum.OrderStatusPending,
			OrderItems: orderItems,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// 재고 차감

		// 카트비우기

	})
}
