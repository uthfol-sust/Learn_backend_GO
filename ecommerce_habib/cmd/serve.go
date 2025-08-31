package cmd

import (
	"ecommerce/controllers"
	"ecommerce/middleware"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	router := http.NewServeMux()

	// Routes
	router.Handle("GET /products", middleware.CorsMiddleware(http.HandlerFunc(controllers.GetProducts)))
	router.Handle("POST /products", middleware.CorsMiddleware(http.HandlerFunc(controllers.CreateProducts)))
    router.Handle("GET /products/{id}",middleware.CorsMiddleware(http.HandlerFunc(controllers.GetProductById)))
	
	fmt.Println("Server running on: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

