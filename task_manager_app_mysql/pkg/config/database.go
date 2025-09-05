package config

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB


func DBConnection() {
	cfg := mysql.NewConfig()

	cfg.User = LocalConfig.DBUser
	cfg.Passwd =LocalConfig.DBPassword
	cfg.Net = LocalConfig.DBNet
	cfg.Addr = LocalConfig.DBHost + ":" + LocalConfig.DBPort
	cfg.DBName = LocalConfig.DBName

	d, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
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