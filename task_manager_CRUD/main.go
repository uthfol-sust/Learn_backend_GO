package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	DueDate       time.Time `json:"due_date"`
	Priority      string    `json:"priority"`
	Category      string    `json:"category"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Notifications bool      `json:"notifications"`
}

var Tasks []Task
var notifi = 24

func main() {
	router := mux.NewRouter()

	// Dummy task
	Tasks = append(Tasks, Task{
		ID:            1,
		Title:         "Finish Golang project",
		Description:   "Complete the REST API using Gorilla Mux",
		Status:        "pending",
		DueDate:       time.Now().Add(time.Duration(notifi) * time.Hour),
		Priority:      "high",
		Category:      "work",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Notifications: true,
	})

	// API routes
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")

	//Serve frontend files
	fs := http.FileServer(http.Dir("./frontend"))
	router.PathPrefix("/").Handler(fs)

	fmt.Println("Server Listening on Port 8000..")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, item := range Tasks {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
