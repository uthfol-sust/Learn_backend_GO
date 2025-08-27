package routes

import(
	"github.com/gorilla/mux"
	"taskmanager/pkg/controllers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/signup",controllers.SignUp).Methods("POST")
	router.HandleFunc("/login",controllers.Login).Methods("POST")


	router.HandleFunc("/users",controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",controllers.DeleteUser).Methods("DELETE")
}