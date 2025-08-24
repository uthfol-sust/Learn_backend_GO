package config

import(
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

var(
	db * gorm.DB
)

func Connect(){
	Str := os.Getenv("DB_String")
	d ,err := gorm.Open("mysql", Str )

	if err !=nil{
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB{
	return db
}

