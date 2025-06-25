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
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/me", controllers.GetMe)
		auth.GET("/cart", controllers.GetCartByUserID)
		auth.POST("/cart", controllers.AddToCart)

	}

	admin := r.Group("/admin")
	admin.Use(middlewares.JWTAuthMiddleware())
	admin.Use(middlewares.AdminOnlyMiddleware())
	{
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)
	}

	return r
}
