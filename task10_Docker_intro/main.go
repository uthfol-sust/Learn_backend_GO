package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	// extract name from query param: /Hello?name=John
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Invalid Request: missing 'name'", http.StatusBadRequest)
		return
	}

	// prepare JSON response
	response := map[string]string{
		"message": name + " hello",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := http.NewServeMux()

	// Example: /Hello?name=John
	router.Handle("GET /Hello", http.HandlerFunc(handlerHello))

	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
