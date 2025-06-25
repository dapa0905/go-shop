package services

import (
	"go-shop/config"
	"go-shop/dtos"
	"go-shop/models"

	"gorm.io/gorm"
)

func GetCartByUserID(userID uint) ([]dtos.CartItemResponse, error) {
	var cart models.Cart

	err := config.DB.Preload("CartItems.Product").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []dtos.CartItemResponse{}, nil
		}
		return nil, err
	}

	var dtoItems []dtos.CartItemResponse
	for _, item := range cart.CartItems {
		dtoItems = append(dtoItems, dtos.CartItemResponse{
			ItemID:      item.ID,
			ProductID:   item.Product.ID,
			ProductName: item.Product.Name,
			Price:       item.Product.Price,
			Quantity:    item.Quantity,
		})
	}
	return dtoItems, nil

}

func AddToCart(userID uint, req dtos.AddToCartRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var cart models.Cart
		if err := tx.Where("user_id = ?", userID).First(&cart).Error; err == gorm.ErrRecordNotFound {
			cart = models.Cart{UserID: userID}
			if err := tx.Create(&cart).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		var cartItem models.CartItem
		err := tx.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).
			First(&cartItem).Error

		if err == gorm.ErrRecordNotFound {
			newItem := models.CartItem{CartID: cart.ID, ProductID: req.ProductID, Quantity: req.Quantity}
			return tx.Create(&newItem).Error
		} else if err != nil {
			return err
		}

		cartItem.Quantity += req.Quantity
		return tx.Save(&cartItem).Error
	})
}
