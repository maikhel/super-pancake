package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB
var err error

type Product struct {
	ID             int       `gorm:"primary_key"`
	Name           string    `gorm:"type:varchar(255)"`
	Amount         int       `gorm:"type:int"`
	Weight         float64   `gorm:"type:float(2)"`
	ExpirationDate time.Time `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"default:NOW()"`
	UpdatedAt      time.Time `gorm:"default:NOW()"`
}

var DatabaseHost = os.Getenv("DB_HOST")
var DatabaseUser = os.Getenv("DB_USER")
var DatabaseName = os.Getenv("DB_NAME")
var DatabasePassword = os.Getenv("DB_PASSWORD")
var DatabasePort = os.Getenv("DB_PORT")

func main() {

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s password=%s sslmode=disable connect_timeout=5",
		DatabaseHost, DatabaseUser, DatabaseName, DatabasePort, DatabasePassword,
	)

	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&Product{})

	r := gin.Default()
	r.GET("/products/", GetProducts)
	r.GET("/products/:id", ShowProduct)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.Run(":8080")

}

func CreateProduct(c *gin.Context) {
	var product Product
	c.BindJSON(&product)

	if err := db.Create(&product).Error; err != nil {
		c.AbortWithStatus(422)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&product)
	db.Save(&product)

	c.JSON(200, product)
}

func ShowProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	var product Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)

		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
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

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	var product Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Delete(&product)
		c.JSON(200, product)
	}
}
