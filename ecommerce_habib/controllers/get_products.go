package controllers

import (
	"net/http"
	"ecommerce/utils"
	"ecommerce/database"

)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductsList, http.StatusOK)
}
