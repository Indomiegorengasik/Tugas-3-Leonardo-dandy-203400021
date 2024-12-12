package main

import (
	"fmt"
	"jualhp/handlers" // Correct import for the handlers package
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
