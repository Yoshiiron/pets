package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("error loading envs from .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, db_name)
	fmt.Println(dsn)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных: " + err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
