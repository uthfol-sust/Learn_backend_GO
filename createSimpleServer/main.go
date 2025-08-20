package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Movie struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

var movies []Movie

func main() {

	router := mux.NewRouter()

	movies = append(movies, Movie{ID: 1,
		Name: "Inception",
		Isbn: "12345",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		}})
	movies = append(movies, Movie{ID: 2,
		Name: "Inception",
		Isbn: "12345",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies",createMovie).Methods("POST")
    router.HandleFunc("/movies",updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Server Listening on port 8000..\n")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// r.HandleFunc("/movies",createMovie).Methods("POST")
// r.HandleFunc("/movies",updateMovie).Methods("PUT")
// r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")


func createMovie(w http.ResponseWriter, r *http.Request) {

	 if r.Header.Get("Content-Type")!="application/json"{
        writeJSONError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
        return
	 }

     w.Header().Set("Content-Type","application/json")

	 var movie Movie
	 json.NewDecoder(r.Body).Decode(&movie)

	 
     movie.ID = rand.Intn(1000000)
	
	 movies = append(movies, movie)

	 json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type")!="application/json"{
      writeJSONError(w,http.StatusUnsupportedMediaType,"Content-Type must be application/json")
	  return
	}

	

	var movie Movie
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
    writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
    return
	}

    fmt.Print(movie.ID)
	
	for index ,item := range movies{
		if(item.ID==movie.ID){

			fmt.Printf("Before update: %+v\n", item)

			if movie.Name != "" {
				item.Name = movie.Name
			}
			if movie.Isbn != "" {
				item.Isbn = movie.Isbn
			}
			if movie.Director != nil {
				if movie.Director.FirstName != "" {
					if item.Director == nil {
						item.Director = &Director{} // allocate if missing
					}
					item.Director.FirstName = movie.Director.FirstName
				}
				if movie.Director.LastName != "" {
					if item.Director == nil {
						item.Director = &Director{}
					}
					item.Director.LastName = movie.Director.LastName
				}
            }

			fmt.Printf("After update: %+v\n", item)
			movies[index] = item
		     
			json.NewEncoder(w).Encode(item)
			return
		}
	}


}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
		    break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)

}

func writeJSONError(w http.ResponseWriter, status int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    json.NewEncoder(w).Encode(map[string]string{
        "error":   message,
        "status":  http.StatusText(status),
    })
}
