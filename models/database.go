package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

var DatabaseHost = os.Getenv("DB_HOST")
var DatabaseUser = os.Getenv("DB_USER")
var DatabaseName = os.Getenv("DB_NAME")
var DatabasePassword = os.Getenv("DB_PASSWORD")
var DatabasePort = os.Getenv("DB_PORT")

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s password=%s sslmode=disable connect_timeout=5",
		DatabaseHost, DatabaseUser, DatabaseName, DatabasePort, DatabasePassword,
	)
	fmt.Printf(dsn)

	DB, err = gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	DB.AutoMigrate(&Product{})
}
