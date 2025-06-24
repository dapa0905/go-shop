package services

import (
	"go-shop/dtos"

	"github.com/gin-gonic/gin"
)

func GetAllCart(c *gin.Context) ([]dtos.GetCartResponse, error) {
	var userId = c.Get("userID")
	if userId = nil {
		
	}

}
