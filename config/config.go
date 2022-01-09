package config

import (
	"os"
	"restful-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config := os.Getenv("CONNECTION_DB")
	var e error

	DB, e := gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate(DB)
	return DB
}

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
