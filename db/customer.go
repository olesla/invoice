package db

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        uint
	Invoices  []Invoice
	Name      string
	Number    string
	Address   string
	Email     string
	Org       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetCustomer(id string) Customer {
	var customer Customer
	DB.First(&customer, id)
	return customer
}

func GetCustomerInvoices(id string) Customer {
	var customer Customer
	DB.Model(&Customer{}).Where("id = ?", id).Preload("Invoices", func(db *gorm.DB) *gorm.DB {
		return db.Order("invoices.id DESC")
	}).Find(&customer)
	return customer
}

func GetCustomers() []Customer {
	var customers []Customer
	DB.Find(&customers)
	return customers
}
