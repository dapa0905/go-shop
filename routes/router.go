package routes

import (
	"go-shop/controllers"
	"go-shop/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/me", controllers.GetMe)
		auth.POST("/products", controllers.CreateProduct)
	}

	return r
}
