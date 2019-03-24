package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/maikhel/food-app/models"
)

var Environment = os.Getenv("ENVIRONMENT")

func main() {

	fmt.Printf("ENVIRONMENT SET TO=%+v\n", Environment)

	if Environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	r.GET("/products/", GetProducts)
	r.GET("/products/:id", ShowProduct)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.Run(":" + os.Getenv("PORT"))

}

func CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)

	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatus(422)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&product)
	models.DB.Save(&product)

	c.JSON(200, product)
}

func ShowProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	var product models.Product
	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)

		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, map[string]interface{}{"products": products})
	}

}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	var product models.Product
	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		models.DB.Delete(&product)
		c.JSON(200, product)
	}
}
