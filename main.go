package main

import (
	"encoding/json"
	"learnGo/config"
	"learnGo/database"
	"learnGo/handlers"
	"learnGo/repositories"
	"learnGo/services"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Categories struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"deskription"`
}

var categories = []Categories{
	{ID: 1, Nama: "Food", Description: "Food Description"},
	{ID: 2, Nama: "Drink", Description: "Drink Description"},
	{ID: 3, Nama: "Beverages", Description: "Food Description"},
}

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapsctructure:"DB_CONN"`
}

func main() {
	log.Println("Program started 🚀")

	// Load config (works for local + Zeabur)
	cfg := config.LoadConfig()

	// Init database
	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/product", productHandler.HandleProducts)
	mux.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Server config (important for production)
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server running on port", cfg.Port)
	log.Fatal(server.ListenAndServe())
}

func getCategoriesByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid categories ID", http.StatusBadRequest)
		return
	}

	// Cari categories dengan ID tersebut
	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Categories belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/categories/{id}
func updateCategories(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Categories ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCategories Categories
	err = json.NewDecoder(r.Body).Decode(&updateCategories)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop categories, cari id, ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updateCategories.ID = id
			categories[i] = updateCategories

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategories)
			return
		}
	}

	http.Error(w, "Categories belum ada", http.StatusNotFound)
}

func deleteCategories(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Categories ID", http.StatusBadRequest)
		return
	}

	// loop categories cari ID, dapet index yang mau dihapus
	for i, p := range categories {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Categories belum ada", http.StatusNotFound)
}
