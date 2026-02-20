package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file not found")
	}

	expString, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Name:     os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			TZ:       os.Getenv("DB_TZ"),
		},
		JWT: JWT{
			Key: os.Getenv("JWT_KEY"),
			Exp: expString,
		},
	}

}
