package controllers

import (
	"go-shop/dtos"
	"go-shop/services"
	"net/http"
	"strconv"

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

func UpdateCartItem(c *gin.Context) {
	var req dtos.UpdateToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uid.(uint)

	err := services.UpdateCartItem(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shopping cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "A product has been updated to your shopping cart."})
}

func DeleteCartItem(c *gin.Context) {
	idParam := c.Param("product_id")
	productID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uid.(uint)

	err = services.DeleteCartItem(userID, uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to delete shopping cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted."})

}

func ClearCartItem(c *gin.Context) {
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uid.(uint)
	err := services.ClearCartItem(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to delete shpping cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item clear."})

}
