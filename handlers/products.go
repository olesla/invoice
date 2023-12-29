package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"invoice/db"
	"net/http"
	"path/filepath"
)

func ProductIndex(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
		Products  []db.Product
	}
	products := db.GetProducts()
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "product", "index.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Produkter",
		Products:  products,
	})
}

func ProductCreate(w http.ResponseWriter, req *http.Request) {
	type PageData struct {
		PageTitle string
	}
	base := filepath.Join("html", "base.html")
	file := filepath.Join("html", "product", "create.html")
	tmpl := template.Must(template.ParseFiles(base, file))
	tmpl.Execute(w, PageData{
		PageTitle: "Opprett",
	})
}

func ProductSave(w http.ResponseWriter, req *http.Request) {
	var product db.Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&product)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	result := db.DB.Create(&product)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(product)
	w.Write(response)
}
