package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Getenv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading enviroment variable %v", err)
	}
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
