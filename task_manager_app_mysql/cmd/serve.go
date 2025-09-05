package cmd

import (
	"fmt"
	"log"
	"net/http"
	"taskmanager/pkg/config"
	"taskmanager/pkg/models"
	"taskmanager/pkg/routes"

	"github.com/joho/godotenv"
)

func Serve() {
	errENV := godotenv.Load()

	if errENV != nil {
		log.Fatal("Error loading .env file")
	}

	//database connection
	config.Connection()
	models.UserAutoMigrate()
	models.TaskAutoMigrate()
	models.EmailVerificationAutoMigrate()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	fmt.Print("Server Running on port 8000..\n")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
