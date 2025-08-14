package main

import (

	"everything-verse/internal/middleware"
	"everything-verse/internal/routes"
	"everything-verse/jobs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "everything-verse/database"
	_ "everything-verse/jobs"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Cors)
	app.Use(middleware.Logger)
	app.Use(middleware.Compress)
	app.Use(middleware.Security)
	app.Use(middleware.RateLimit)
	app.Use(recover.New())

	service := app.Group("/")

	routes.Api(service)
	go jobs.Cron()

	app.Use(middleware.NotFound)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
