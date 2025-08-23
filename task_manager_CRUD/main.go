package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
    "math/rand" 

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/tasks",createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
    router.HandleFunc("/tasks/{id}",deleteTask).Methods("DELETE")

	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","PUT","PATCH","DELETE","OPTIONS"},
		AllowedHeaders:  []string{"Content-Type"},
        AllowCredentials: true,
	}

	cores := cors.New(corsOptions)
	handler := cores.Handler(router)

	fmt.Println("Server Listening on Port 8000..")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func createTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
    
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err!=nil{
		http.Error(w,"Invaild JSON format",http.StatusBadRequest)
		return
	}
    
	task.ID = rand.Intn(10000)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	Tasks = append(Tasks, task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
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

func deleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	id ,err := strconv.Atoi(params["id"])
	
	if err != nil{
       http.Error(w,"Invalid ID",http.StatusBadRequest)
	   return
	}

	for index , item := range Tasks{
		if item.ID == id{
            Tasks = append(Tasks[:index],Tasks[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Tasks)
}
