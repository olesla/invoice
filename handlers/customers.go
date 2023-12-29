package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"invoice/db"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func CustomerIndex(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Customers []db.Customer
	}
	customers := db.GetCustomers()
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "customer", "index.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Kunder",
		Customers: customers,
	})
}

func CustomerShow(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Customer  db.Customer
	}
	vars := mux.Vars(req)
	customer := db.GetCustomer(vars["id"])
	invoices := db.GetCustomerInvoices(vars["id"])
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "customer", "show.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: customer.Name,
		Customer:  invoices,
	})
}

func CustomerCreate(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
	}
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "customer", "create.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Opprett",
	})
}

func CustomerSave(w http.ResponseWriter, req *http.Request) {
	var customer db.Customer
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&customer)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	result := db.DB.Create(&customer)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(customer)
	w.Write(response)
}

func CustomerUpdate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var values map[string]string
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&values)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	var customer db.Customer
	result := db.DB.First(&customer, id)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	if values["name"] != "" {
		customer.Name = values["name"]
	}
	if values["number"] != "" {
		customer.Number = values["number"]
	}
	if values["address"] != "" {
		customer.Address = values["address"]
	}
	if values["email"] != "" {
		customer.Email = values["email"]
	}

	db.DB.Save(&customer)
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
