package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskmanager/pkg/models"
	"taskmanager/pkg/utils"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	createTask := &models.Task{}

	utils.ParseBody(r, createTask)

	savedTask, err := models.CreateTask(createTask)

	if err != nil {
		http.Error(w, "Failed to Create New task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedTask)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	var err error

	tasks, err = models.GetAllTasks()

	if err != nil {
		http.Error(w, "Failed to fetch All users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	req_id, err := strconv.Atoi(idstr)
	if err != nil {
		utils.ThrowError(w, "Task ID is Missing at Histroy", 404)
		return
	}

	task := &models.Task{}

	task, err = models.GetTaskByID(req_id)
	if err != nil {
		http.Error(w, "Failed to load this task", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	req_id, _ := strconv.Atoi(idstr)

	updateinfo := &models.Task{}
	utils.ParseBody(r, updateinfo)

	load_task := &models.Task{}
	load_task, _ = models.GetTaskByID(req_id)

	if updateinfo.Title != "" {
		load_task.Title = updateinfo.Title
	}
	if updateinfo.Description != "" {
		load_task.Description = updateinfo.Description
	}
	if updateinfo.Status != "" {
		load_task.Status = updateinfo.Status
	}

	saved, err := models.UpdateTask(load_task)
	if err != nil {
		http.Error(w, "update unsuccessful", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(saved)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	req_id, _ := strconv.Atoi(idstr)

	if err := models.DeleteTask(req_id); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"User deleted successfully"}`))
}
