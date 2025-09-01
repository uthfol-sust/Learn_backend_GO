package cmd

import (
	"ecommerce/middleware"
	"ecommerce/routes"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
    middlewareManager := middleware.Manager{}
	middlewareManager.Use(middleware.Test, middleware.MiddleTest)

	router := http.NewServeMux()

	routes.InitialRoutes(router, &middlewareManager)

	fmt.Println("Server running on: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
