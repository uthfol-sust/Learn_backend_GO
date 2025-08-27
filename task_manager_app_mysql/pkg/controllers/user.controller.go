package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/pkg/models"
	"taskmanager/pkg/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	// hash password
	hashValue, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password: "+err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashValue

	// save user
	savedUser, err := models.UserRegistration(user)
	if err != nil {
		http.Error(w, "Failed to register user to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedUser)
}


func Login(w http.ResponseWriter, r * http.Request){
	
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){

	var users []models.User

    users , err := models.GetAllUsers()

	if err!=nil{
		http.Error(w,"Failed to load users from database",http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}