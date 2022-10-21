package database

import (
	"fmt"
	"log"
	"simple-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "simpleapp"
)

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	DB.AutoMigrate(&models.Book{})

	log.Println("connected to the database")
}
