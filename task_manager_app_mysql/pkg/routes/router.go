package routes

import(
	"github.com/gorilla/mux"
	"taskmanager/pkg/controllers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/signup",controllers.SignUp).Methods("POST")
	router.HandleFunc("/login",controllers.Login).Methods("POST")


	router.HandleFunc("/tasks",controllers.GetAllUsers).Methods("GET")
}