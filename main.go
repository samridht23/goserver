package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/samridht23/server/database"
	"github.com/samridht23/server/router"
	"log"
	"os"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	if err := database.ConnectDb(); err != nil {
		log.Fatal(err)
	}
	godotenv.Load()
	app := fiber.New()
	router.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
