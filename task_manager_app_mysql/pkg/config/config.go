package config

import (
	"os"
)

type Config struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBNet      string `json:"db_net"`
	JWTSecret  string `json:"jwt_secret"`
}

var LocalConfig *Config

func initConfig() *Config {
	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBNet: os.Getenv("DB_Net"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

func SetConfig() {
	LocalConfig = initConfig()
}
