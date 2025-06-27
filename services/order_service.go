package services

import (
	"go-shop/config"
	"go-shop/dtos"
	"go-shop/models"

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

		// Order 생성 및 저장
		// 재고 차감
		// 카트비우기

	})
}
