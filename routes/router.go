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
		auth.GET("/products", controllers.GetAllProducts)
		auth.GET("/products/:id", controllers.GetProductByID)
		auth.PUT("/products/:id", controllers.UpdateProduct)
	}

	return r
}
