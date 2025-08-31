package controllers

import (
	"encoding/json"
	"net/http"
	"ecommerce/utils"
	"ecommerce/database"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Products
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Plz Give Me Vaild Json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductsList) + 1
	database.ProductsList = append(database.ProductsList, newProduct)
	utils.SendData(w, newProduct, http.StatusCreated)
}
