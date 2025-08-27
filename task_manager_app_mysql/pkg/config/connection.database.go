package config

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB


func Connection(){
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pass
	cfg.Net ="tcp"
	cfg.Addr="127.0.0.1:3306"
	cfg.DBName = dbName


	  d , err := sql.Open("mysql",cfg.FormatDSN())

	  if err!=nil{
		log.Fatal(err)
	  }

	  if err := d.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")

	db = d
}

func  GetDB() *sql.DB{
	return  db 
}