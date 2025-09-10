package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"taskmanager/pkg/models"
	"taskmanager/pkg/utils"
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
	if savedUser == nil {
		http.Error(w, "Failed to register user: got nil user", http.StatusInternalServerError)
		return
	}

	code, errVerify := utils.SendVerificationEmail(savedUser.Email)
	if errVerify != nil {
		http.Error(w, "Failed to send verification code: "+errVerify.Error(), http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(5 * time.Minute)
	errSave := models.SaveVerification(savedUser.UserID, savedUser.Email, code, expiresAt)
	if errSave != nil {
		http.Error(w, "Failed to save verification code: "+errSave.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Verification code sent to " + savedUser.Email,
	})
}

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &LoginInput{}
	utils.ParseBody(r, user)

	load_user, err := models.FindUserByEmail(user.Email)

	if err != nil {
		http.Error(w, "Using Unvaild Email to login",  http.StatusUnauthorized)
		return
	}

	verification, err := models.GetVerificationByEmail(user.Email)
	if err != nil || verification == nil {
		http.Error(w, "No user verification record found", http.StatusUnauthorized)
		return
	}
	if !verification.IsVerified {
		http.Error(w, "Your Email is not verified", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(load_user.Password, user.Password) {
		http.Error(w, "Wrong Password", http.StatusNonAuthoritativeInfo)
		return
	}

	token, er := utils.GenerateToken(load_user.UserID, load_user.Email, load_user.Role)
	if er != nil {
		http.Error(w, "Failed to Gernerate Token!", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"massage": "\nLogin successful! JWT set in cookie",
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	users, err := models.GetAllUsers()

	if err != nil {
		http.Error(w, "Failed to load users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		utils.ThrowError(w, "Missing user ID", 404)
		return
	}

	user_id, _ := strconv.Atoi(idStr)

	user, err := models.GetUserByID(user_id)

	if err != nil {
		utils.ThrowError(w, "User Not Exist", 500)
		return
	}

	json_user, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(json_user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	idstr := r.PathValue("id")

	user_id, _ := strconv.Atoi(idstr)

	user, err := models.GetUserByID(user_id)
	if err != nil {
		http.Error(w, "Failed to load users", http.StatusInternalServerError)
		return
	}

	updateUser := &models.User{}
	utils.ParseBody(r, updateUser)

	var errhash error
	updateUser.Password, errhash = utils.HashPassword(updateUser.Password)
	if errhash != nil {
		fmt.Print("New password hashing Error!")
		return
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}

	if updateUser.Password != "" {
		user.Password = updateUser.Password
	}
	if updateUser.Role != "" {
		user.Role = updateUser.Role
	}

	err_saved := models.UpdateUser(user)

	if err_saved != nil {
		http.Error(w, "New Data Not Updated!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)

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

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	code := r.URL.Query().Get("code")

	if email == "" || code == "" {
		http.Error(w, "Missing email or code", http.StatusBadRequest)
		return
	}

	verification, err := models.GetVerificationByEmail(email)

	fmt.Print(err)

	if err != nil {
		http.Error(w, "No verification record found for this email", http.StatusBadRequest)
		return
	}

	if verification.IsVerified {
		http.Error(w, "Email already verified", http.StatusBadRequest)
		return
	}

	if time.Now().After(verification.ExpiresAt) {
		http.Error(w, "Verification code expired", http.StatusUnauthorized)
		return
	}

	err = models.MarkVerified(email)
	if err != nil {
		http.Error(w, "Failed to update verification status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Email verified successfully"}`))
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	ID, _ := strconv.Atoi(idstr)

	var reset = &models.Reset{}

	utils.ParseBody(r, reset)
	reset.UserID = ID

	user, err := models.GetUserByID(ID)
	if err != nil {
		utils.ThrowError(w, "User not found", 404)
	}

	code, msg := utils.SendVerificationEmail(user.Email)
	if msg != nil {
		utils.ThrowError(w, "Failed to send Code", 500)
		return
	}
	reset.Sendcode, _ = strconv.Atoi(code)
	hashPass, err := utils.HashPassword(reset.NewPassword)
	if err != nil {
		utils.ThrowError(w, "New Password can't Hashing", 400)
		return
	}
	reset.NewPassword = hashPass

	if err := models.SaveResetCode(reset); err != nil {
		utils.ThrowError(w, "Unable to save reset code", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Verification code sent. Please check your email.",
	})

}

func ResetCodeVerification(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	strCode := r.URL.Query().Get("code")
	code, _ := strconv.Atoi(strCode)
	id, err := strconv.Atoi(strId)

	orginalCode, newPassword, err := models.GetResetCode(id)
	if err != nil {
		utils.ThrowError(w, "Unable to load reset code", http.StatusInternalServerError)
		return
	}

	if orginalCode != code {
		utils.ThrowError(w, "Invalid verification code", http.StatusUnauthorized)
		return
	}

	user, err := models.GetUserByID(id)
	if err != nil {
		utils.ThrowError(w, "User not found for update Password", http.StatusNotFound)
		return
	}

	user.Password = newPassword
	if err := models.UpdateUser(user); err != nil {
		utils.ThrowError(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	if err := models.DeleteResetCode(id); err != nil {
    utils.ThrowError(w, "Failed to delete reset code", 500)
    return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Password reset successful",
	})
}
