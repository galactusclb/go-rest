package product

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"

)

type Product struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	CategoryID string `json:"category_id"`
}

var products = []Product{
	{ID: "1", Name: "Smartphone", CategoryID: "1"},
	{ID: "2", Name: "Laptop", CategoryID: "1"},
}

func GetProducts(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]

	product := Product{ID: id, Name: "Test", CategoryID: "2"}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var newProduct Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	newProduct.ID = fmt.Sprintf("%d", len(products) + 1)
	log.Println(len(products))

	products = append(products, newProduct)
	json.NewEncoder(w).Encode(newProduct)
}