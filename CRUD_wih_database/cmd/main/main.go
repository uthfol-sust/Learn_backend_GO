package main

import (
	"fmt"
	"log"
	"net/http"
     "os"

	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/routes"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/config"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	config.Connect()
    models.Migrate()

	fmt.Println("DB_String:", os.Getenv("DB_String"))

	router := mux.NewRouter()

	routes.BooksRoutes(router)

	// http.Handle("/",router)

	fmt.Print("Server Running on Port 8000..\n")

	log.Fatal(http.ListenAndServe(":8000",router))
}