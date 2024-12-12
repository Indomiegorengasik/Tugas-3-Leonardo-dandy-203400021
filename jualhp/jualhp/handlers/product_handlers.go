package handlers

import (
	"encoding/json"
	"jualhp/models" // Correct import for the models package
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	products = make(map[string]models.Product)
	mu       sync.Mutex
)

// CreateProduct handles POST requests to create a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	products[product.ID] = product
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProduct handles GET requests to fetch a product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	product, exists := products[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// GetAllProducts handles GET requests to fetch all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	productList := []models.Product{}
	for _, product := range products {
		productList = append(productList, product)
	}

	json.NewEncoder(w).Encode(productList)
}

// UpdateProduct handles PUT requests to update an existing product by ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedProduct models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Check if the product exists
	product, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Update the product
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	// Update any other fields as necessary

	products[id] = product

	// Respond with the updated product
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct handles DELETE requests to delete a product by ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	defer mu.Unlock()

	// Check if the product exists
	if _, exists := products[id]; !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Delete the product
	delete(products, id)

	// Respond with success message
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
