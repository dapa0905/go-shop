package controllers

import (
	"go-shop/dtos"
	"go-shop/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var req dtos.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.RegisterUser(req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "회원가입 실패"})
		return
	}

	res := dtos.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	c.JSON(http.StatusOK, res)
}
