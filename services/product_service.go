package services

import (
	"go-shop/config"
	"go-shop/dtos"
	"go-shop/models"
)

func CreateProduct(req dtos.CreateProductRequest) (*models.Product, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
