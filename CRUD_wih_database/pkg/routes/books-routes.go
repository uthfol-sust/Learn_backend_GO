package routes

import (

	"github.com/gorilla/mux"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/controllers"
)

var BooksRoutes = func(router *mux.Router) { 
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
