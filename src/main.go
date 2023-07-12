package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oatsaysai/simple-core-bank/src/routes"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.DefaultRouter(api)

	app.Listen(":8080")
}
