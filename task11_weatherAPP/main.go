package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ConfigKey struct {
	OpenWeatherAPI string `json:"OpenWeatherAPI"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func loadAPIkey() (ConfigKey, error) {
	data, err := os.ReadFile(".configAPI")
	if err != nil {
		return ConfigKey{}, fmt.Errorf("error reading file: %w", err)
	}

	var configAPI ConfigKey
	if err := json.Unmarshal(data, &configAPI); err != nil {
		return ConfigKey{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return configAPI, nil
}

func FindWeather(w http.ResponseWriter, r *http.Request) {
	apiConfig, err := loadAPIkey()
	if err != nil {
		http.Error(w, "Unable to load API key", http.StatusInternalServerError)
		return
	}

	// Example: city name via query parameter, default = London
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "London"
	}

	// Build API request
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiConfig.OpenWeatherAPI)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to call weather API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read weather API response", http.StatusInternalServerError)
		return
	}

	// Decode JSON
	var weather weatherData
	if err := json.Unmarshal(body, &weather); err != nil {
		http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
		return
	}

	// Send JSON response to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func main() {
	router := http.NewServeMux()
	router.Handle("GET /weather", http.HandlerFunc(FindWeather))

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
