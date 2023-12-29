package db

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Code        string
	Price       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetProduct(id string) Product {
	var product Product
	DB.First(&product, id)
	return product
}

func GetInvoiceProducts(id string) Invoice {
	var invoice Invoice
	DB.Model(&Invoice{}).Preload("Products").Find(&invoice)
	return invoice
}

func GetProducts() []Product {
	var products []Product
	DB.Find(&products)
	return products
}
