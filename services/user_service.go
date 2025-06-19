package services

import (
	"errors"
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

func LoginUser(email, password string) (*models.User, error) {

	var user models.User

	if err := config.DB.Where("email =?", email).First(&user).Error; err != nil {
		return nil, errors.New("not find User")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("wrong Password")
	}

	return &user, nil

}
