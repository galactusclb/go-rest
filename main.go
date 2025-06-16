package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"CLB-go-rest/product"

)

func handler(w http.ResponseWriter, r *http.Request){
	// http.handler('/')
	fmt.Fprintf(w, "hello Test 3\n")
}

func main(){
	
	r := mux.NewRouter()
 
	r.HandleFunc("/", handler)
	api := r.PathPrefix("/api/v1").Subrouter()

	productRouter := api.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", product.GetProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", product.GetProduct).Methods(http.MethodGet)
	productRouter.HandleFunc("", product.CreateProduct).Methods(http.MethodPost)

	log.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", r))
}