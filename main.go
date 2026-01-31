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
	"time"
)

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

	// Dependency Injection
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/product", productHandler.HandleProducts)
	mux.HandleFunc("/api/product/", productHandler.HandleProductByID)

	mux.HandleFunc("/api/category", categoryHandler.HandleCategories)
	mux.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

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
