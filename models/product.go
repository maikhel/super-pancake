package models

import (
	"time"
)

type Product struct {
	ID             int       `gorm:"primary_key"`
	Name           string    `gorm:"type:varchar(255)"`
	Amount         int       `gorm:"type:int"`
	Weight         float64   `gorm:"type:float(2)"`
	ExpirationDate time.Time `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"default:NOW()"`
	UpdatedAt      time.Time `gorm:"default:NOW()"`
}

func GetProducts() (*[]Product, error) {
	var products []Product
	err := DB.Find(&products).Error

	return &products, err
}

func GetProduct(id int) (*Product, error) {
	var product Product

	err := DB.Where("id = ?", id).First(&product).Error

	return &product, err
}

func CreateProduct(input *Product) (*Product, error) {
	err := DB.Create(&input).Error

	return input, err
}

func UpdateProduct(product *Product, payload map[string]interface{}) (*Product, error) {
	err := DB.Model(&product).Updates(payload).Error

	return product, err
}

func DeleteProduct(product *Product) (*Product, error) {
	err := DB.Delete(&product).Error

	return product, err
}
