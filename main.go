package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Product struct {
	ID             int       `gorm:"primary_key"`
	Name           string    `gorm:"type:varchar(255)"`
	Amount         int       `gorm:"type:int"`
	Weight         float64   `gorm:"type:float(2)"`
	ExpirationDate time.Time `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"default:NOW()"`
	UpdatedAt      time.Time `gorm:"default:NOW()"`
	DeletedAt      time.Time `gorm:"default:null"`
}

func main() {

	var err error

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=food_user dbname=food_db password=password sslmode=disable connect_timeout=5")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	r := gin.Default()
	r.GET("/", GetProducts)

	r.Run(":8080")

}

func GetProducts(c *gin.Context) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}

}
