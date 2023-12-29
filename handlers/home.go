package handlers

import (
	"fmt"
	"html/template"
	"invoice/db"
	"net/http"
	"path/filepath"
)

var COMPANY_NAME string
var COMPANY_ADDRESS string
var COMPANY_EMAIL string
var COMPANY_PHONE string
var COMPANY_ACCOUNT_NUMBER string
var COMPANY_ORG_NUMBER string

func Home(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Invoices  []db.Invoice
	}
	invoices := db.GetYearlyInvoices()
	fmt.Println(invoices)
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "home.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Oversikt",
		Invoices:  invoices,
	})
}
