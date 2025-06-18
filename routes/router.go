package routes

import (
	"go-shop/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	return r
}
