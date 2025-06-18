package services

import (
	"go-shop/config"
	"go-shop/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password, name string) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
		Role:     "user",
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
