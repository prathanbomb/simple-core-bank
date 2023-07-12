package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DefaultRouter(app fiber.Router) {
	app.Get("/health", Health())
	app.Get("/hello-world", HelloWorld())
}

func Health() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		_, err := c.WriteString("OK")
		return err
	}
}

func HelloWorld() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	}
}
