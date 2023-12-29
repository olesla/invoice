package db

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID         uint
	CustomerID uint
	Customer   Customer
	Products   string
	Number     string
	Sum        float64
	Sent       bool
	Paid       bool
	Date       time.Time
	Due        time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func GetInvoices() []Invoice {
	var invoices []Invoice
	DB.Preload("Customer").Find(&invoices)
	return invoices
}

func GetInvoice(id string) Invoice {
	var invoice Invoice
	DB.Preload("Customer").First(&invoice, id)
	return invoice
}

func GetYearlyInvoices() []Invoice {
	var invoices []Invoice
	DB.Preload("Customer").Where("date LIKE ?", "2023%").Find(&invoices)
	return invoices
}
