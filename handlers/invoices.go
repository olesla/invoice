package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"invoice/db"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func InvoiceIndex(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Invoices  []db.Invoice
	}
	invoices := db.GetInvoices()
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "invoice", "index.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Fakturaer",
		Invoices:  invoices,
	})
}

func InvoiceCreate(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Customer  db.Customer
		Products  []db.Product
	}
	vars := mux.Vars(req)
	customer := db.GetCustomer(vars["id"])
	products := db.GetProducts()
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "invoice", "create.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Opprett faktura",
		Customer:  customer,
		Products:  products,
	})
}

func InvoiceSave(w http.ResponseWriter, req *http.Request) {
	type Body struct {
		Date     string  `json:"date"`
		Due      string  `json:"due"`
		Sum      float64 `json:"sum"`
		Products []struct {
			ID     uint    `json:"id"`
			Amount uint    `json:"amount"`
			Name   string  `json:"name"`
			Price  float64 `json:"price"`
		}
	}
	var body Body
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	var invoice db.Invoice
	invoice.CustomerID = uint(id)
	invoice.Sum = body.Sum
	products, _ := json.Marshal(body.Products)
	invoice.Products = string(products)
	invoice.Sent = false
	invoice.Paid = false
	invoice.Due = time.Now()
	invoice.Date = time.Now()
	invoice.Number = ""

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	due, err := time.Parse("2006-01-02", body.Due)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	invoice.Due = due
	invoice.Date = date

	result := db.DB.Create(&invoice)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	invoice.Number = fmt.Sprintf("%d%d", id, invoice.ID)
	db.DB.Save(&invoice)

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(invoice)
	w.Write(response)
}

func InvoiceShow(w http.ResponseWriter, req *http.Request) {
	type Product struct {
		ID     uint    `json:"id"`
		Amount uint    `json:"amount"`
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
	}
	type PageData struct {
		PageTitle            string
		Invoice              db.Invoice
		Products             []Product
		CompanyName          string
		CompanyAddress       string
		CompanyEmail         string
		CompanyPhone         string
		CompanyAccountNumber string
		CompanyOrgNumber     string
	}
	vars := mux.Vars(req)
	invoice := db.GetInvoice(vars["id"])
	var products []Product
	json.Unmarshal([]byte(invoice.Products), &products)

	view := req.URL.Query().Get("view")

	var tmpl *template.Template

	if view == "print" {
		file := filepath.Join("html", "invoice", "print.html")
		tmpl = template.Must(template.ParseFiles(file))
	} else {
		base := filepath.Join("html", "base.html")
		file := filepath.Join("html", "invoice", "show.html")
		tmpl = template.Must(template.ParseFiles(base, file))
	}

	tmpl.Execute(w, PageData{
		PageTitle:            fmt.Sprintf("%s %s", "Faktura", invoice.Number),
		Invoice:              invoice,
		Products:             products,
		CompanyName:          COMPANY_NAME,
		CompanyAddress:       COMPANY_ADDRESS,
		CompanyEmail:         COMPANY_EMAIL,
		CompanyPhone:         COMPANY_PHONE,
		CompanyAccountNumber: COMPANY_ACCOUNT_NUMBER,
		CompanyOrgNumber:     COMPANY_ORG_NUMBER,
	})
}
