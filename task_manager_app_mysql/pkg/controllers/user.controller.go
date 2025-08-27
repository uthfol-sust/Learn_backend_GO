package controllers

import (
	"net/http"
	"encoding/json"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request){
  tasks := []map[string]interface{}{
		{"id": 1, "title": "Learn Go", "completed": false},
		{"id": 2, "title": "Build Task Manager", "completed": true},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}