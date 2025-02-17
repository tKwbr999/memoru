package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-github-account/memoru-backend/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":8080")
}
