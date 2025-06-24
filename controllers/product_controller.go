package controllers

import (
	"go-shop/dtos"
	"go-shop/services"
	"net/http"
	"strconv"

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

func GetAllProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load product list"})
		return
	}

	var response []dtos.ProductResponse
	for _, p := range products {
		response = append(response, dtos.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")

	idUnit, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := services.GetProductByID(uint(idUnit))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)

}

func UpdateProduct(c *gin.Context) {

	idParam := c.Param("id")
	idUnit, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req dtos.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.UpdateProduct(uint(idUnit), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dtos.UpdateProductRequest{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       req.Stock,
	}

	c.JSON(http.StatusOK, response)

}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	idUnit, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = services.DeleteProduct(uint(idUnit))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
