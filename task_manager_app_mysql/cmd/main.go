package main

import (
	"fmt"
	"log"
	"net/http"

	"taskmanager/pkg/config"
	"taskmanager/pkg/middleware"
	"taskmanager/pkg/models"
	"taskmanager/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
    errENV := godotenv.Load("../.env")

	if errENV != nil {
	log.Fatal("Error loading .env file")
	}
   
	//database connection
    config.Connection()
	models.UserAutoMigrate()
	models.TaskAutoMigrate()
	models.EmailVerificationAutoMigrate()

	
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	handler := middleware.CorsMiddleware(router)

	


	fmt.Print("Server Running on port 8000..\n")
	log.Fatal(http.ListenAndServe(":8000",handler))
}