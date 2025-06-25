package controllers

import (
	"go-shop/dtos"
	"go-shop/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCartByUserID(c *gin.Context) {
	userID := c.GetUint("user_id")

	cartItems, err := services.GetCartByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, dtos.GetCartResponse{
		CartItems: cartItems,
	})

}

func AddToCart(c *gin.Context) {
	var req dtos.AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	err := services.AddToCart(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add shopping cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "A product has been added to your shopping cart."})

}
