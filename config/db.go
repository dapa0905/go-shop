package config

import (
	"fmt"
	"go-shop/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env 파일 로드 실패:", err)
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := "127.0.0.1"
	port := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("AutoMigrate 실패:", err)
	}

	DB = database

}
