package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"database/sql"
	"errors"

	"taskmanager/pkg/models"
	"taskmanager/pkg/utils"

	"github.com/gorilla/mux"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	hashValue, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password: "+err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashValue

	savedUser, err := models.UserRegistration(user)
	if err != nil {
		http.Error(w, "Failed to register user to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

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
		http.Error(w,"Failed to load users",http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	user_id , _ := strconv.Atoi(vars["id"])

	user , err := models.GetUserByID(user_id)
	if err!=nil{
		http.Error(w,"Not Exist this User",http.StatusInternalServerError)
	}
    
	json_user , _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)
	w.Write(json_user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	user_id , _ := strconv.Atoi(vars["id"])

	user, err := models.GetUserByID(user_id)
	if err!=nil{
		http.Error(w,"Failed to load users",http.StatusInternalServerError)
	}

	updateUser := &models.User{}
	utils.ParseBody(r , updateUser)


	var errhash error
	updateUser.Password , errhash = utils.HashPassword(updateUser.Password)
	if errhash !=nil{
		fmt.Print("New password hashing Error!")
	}

	if updateUser.Name!=""{
       user.Name=updateUser.Name
	}

	if updateUser.Password!=""{
		user.Password = updateUser.Password
	}

	err_saved := models.UpdateUser(user)

	if err_saved!=nil{
		http.Error(w, "New Data Not Updated!",http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := models.DeleteUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"User deleted successfully"}`))
}
