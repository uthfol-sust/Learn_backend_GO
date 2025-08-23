package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func main() {
	erro := godotenv.Load()
	if erro != nil {
	log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	// Database config
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pass
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = dbName

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database!")

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/albums", getAlbums).Methods("GET")

	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	//get data from mysql server
	rows, err := db.Query("SELECT  title FROM album")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var alb Album
		if err := rows.Scan( &alb.Title); err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Row iteration error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(albums)
}