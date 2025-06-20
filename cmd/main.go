package main

import (
	"go-shop/config"
	"go-shop/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
