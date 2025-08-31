package controllers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("id")

	ID , err := strconv.Atoi(Id)
	if err!=nil{
		http.Error(w, "PLz give me a Vaild Product id",400)
	}

	for _, product := range database.ProductsList{
        if product.ID==ID{
           utils.SendData(w , product, 200)
		   return
		}
	}

	utils.SendData(w , "Data Not Found in List",404)
}
