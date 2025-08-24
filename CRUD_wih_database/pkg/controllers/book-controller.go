package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/models"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	NewBook := models.GetAllBooks()

	res ,_ := json.Marshal(NewBook)

	w.Header().Set("Content-Type","pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)

	fmt.Print("Book created Suceessfully")
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	bookID := vars["bookId"]

	ID , err := strconv.ParseInt(bookID,0,0)

	if err!=nil{
       fmt.Print("error while Parsing")
	}

	bookDetails , _ := models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type","pgklication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
    createBook := &models.Book{}

	utils.ParseBody(r , createBook)

	b := createBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.DeleteBook(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}

	utils.ParseBody(r,updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID , err := strconv.ParseInt(bookId,0,0)

	if err!=nil{
       fmt.Print("error while Parsing")
	}

	bookdetails , db := models.GetBookById(ID)

	if updateBook.Name != ""{
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author !=""{
        bookdetails.Author = updateBook.Author
	}

	if updateBook.Publication !=""{
        bookdetails.Publication = updateBook.Publication
	}
    
	db.Save(&bookdetails)


   res, _ := json.Marshal(bookdetails)
   
   w.Header().Set("Content-Type","pgklication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}