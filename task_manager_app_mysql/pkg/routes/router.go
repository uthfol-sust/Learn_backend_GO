package routes

import(
	"github.com/gorilla/mux"
	"taskmanager/pkg/controllers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/tasks",controllers.GetAllTasks).Methods("GET")
}