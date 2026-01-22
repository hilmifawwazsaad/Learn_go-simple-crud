package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simple-crud/handlers"
)

func main() {
	// Welcome
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	// GET /categories → Ambil semua kategori
	// POST /categories → Tambah kategori
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetAllCategories(w, r)
		} else if r.Method == "POST" {
			handlers.CreateCategory(w, r)
		}
	})

	// GET /categories/{id} → Ambil detail satu kategori
	// PUT /categories/{id} → Update kategori
	// DELETE /categories/{id} → Hapus kategori
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetCategoryByID(w, r)
		} else if r.Method == "PUT" {
			handlers.UpdateCategory(w, r)
		} else if r.Method == "DELETE" {
			handlers.DeleteCategory(w, r)
		}
	})

	fmt.Println("Server running di localhost:8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Gagal Running Server")
	}
}
