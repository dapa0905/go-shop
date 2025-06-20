package controllers

import (
	"go-shop/dtos"
	"go-shop/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req dtos.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          product.ID,
		"name":        product.Name,
		"Description": product.Description,
		"Price":       product.Price,
		"stock":       product.Stock,
	})
}
