package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samridht23/server/pkg/router"
	"github.com/samridht23/server/pkg/utils"
	"log"
)

func main() {
	app := fiber.New()
	router.Initalize(app)
	log.Fatal(app.Listen(":" + utils.Getenv("PORT", "3000")))
}
