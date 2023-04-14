package main

import (
	"time"
	"vlab/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Api rate limiter
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
		})
	})

	routes.BasicRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
