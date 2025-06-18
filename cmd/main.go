package main

import (
	"go-shop/config"
	"go-shop/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
