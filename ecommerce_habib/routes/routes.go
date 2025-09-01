package routes

import (
	"ecommerce/controllers"
	"ecommerce/middleware"
	"net/http"
)

func InitialRoutes(router *http.ServeMux ,middlewareManager *middleware.Manager ) {
	router.Handle("GET /products", 
	middlewareManager.With(
		http.HandlerFunc(controllers.GetProducts),
		 middleware.CorsMiddleware),
	)
	// Routes
	// router.Handle("GET /products", middleware.CorsMiddleware(http.HandlerFunc(controllers.GetProducts)))
	router.Handle("POST /products", middleware.CorsMiddleware(http.HandlerFunc(controllers.CreateProducts)))
	router.Handle("GET /products/{id}", middleware.CorsMiddleware(http.HandlerFunc(controllers.GetProductById)))
}
