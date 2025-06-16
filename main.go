package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	// http.handler('/')
	fmt.Fprintf(w, "hello Test\n")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}