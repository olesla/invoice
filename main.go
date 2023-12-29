package main

import (
	"flag"
	"fmt"
	"os"

	"invoice/db"
	"invoice/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	db.DSN = os.Getenv("DSN")
	handlers.COMPANY_NAME = os.Getenv("COMPANY_NAME")
	handlers.COMPANY_ADDRESS = os.Getenv("COMPANY_ADDRESS")
	handlers.COMPANY_EMAIL = os.Getenv("COMPANY_EMAIL")
	handlers.COMPANY_PHONE = os.Getenv("COMPANY_PHONE")
	handlers.COMPANY_ACCOUNT_NUMBER = os.Getenv("COMPANY_ACCOUNT_NUMBER")
	handlers.COMPANY_ORG_NUMBER = os.Getenv("COMPANY_ORG_NUMBER")

	db.StartDb()

	seed := flag.Bool("seed", false, "Seed database with example data")
	flag.Parse()
	if *seed {
		db.Seed()
		fmt.Println("Database successfully seeded")
		os.Exit(0)
	}
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", handlers.Home).Methods("GET")

	r.HandleFunc("/customers", handlers.CustomerIndex).Methods("GET")
	r.HandleFunc("/customers", handlers.CustomerSave).Methods("POST")
	r.HandleFunc("/customers/new", handlers.CustomerCreate).Methods("GET")
	r.HandleFunc("/customers/{id}", handlers.CustomerShow).Methods("GET")
	r.HandleFunc("/customers/{id}", handlers.CustomerUpdate).Methods("PUT")
	r.HandleFunc("/customers/{id}/invoices/new", handlers.InvoiceCreate).Methods("GET")
	r.HandleFunc("/customers/{id}/invoices/create", handlers.InvoiceSave).Methods("POST")

	r.HandleFunc("/products", handlers.ProductIndex).Methods("GET")
	r.HandleFunc("/products/new", handlers.ProductCreate).Methods("GET")
	r.HandleFunc("/products", handlers.ProductSave).Methods("POST")

	r.HandleFunc("/invoices", handlers.InvoiceIndex).Methods("GET")
	r.HandleFunc("/invoices/{id}", handlers.InvoiceShow).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Starting at http://localhost:8090")
	http.ListenAndServe(":8090", r)
}
