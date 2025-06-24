package services

import (
	"fmt"
	"go-shop/config"
	"go-shop/dtos"
	"go-shop/models"

	"gorm.io/gorm"
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

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("can't find the product")
		}

		return nil, err
	}
	return &product, nil

}

func UpdateProduct(id uint, dto dtos.UpdateProductRequest) (*models.Product, error) {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("can't find the product")
		}
		return nil, err
	}

	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Stock = dto.Stock

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil

}

func DeleteProduct(id uint) error {
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("can't find the product")
		}
		return err
	}

	result := config.DB.Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("delete failed: no rows affected")
	}

	return nil
}
