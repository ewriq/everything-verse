package routes

import (
	"everything-verse/internal/handler"
	"everything-verse/internal/middleware"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3"
)

func Api(app fiber.Router) {
	app.Use(middleware.Cors)
	app.Use(middleware.Logger)
	app.Use(middleware.Compress)
	app.Use(middleware.Security)
	app.Use(middleware.RateLimit)
	app.Use(recover.New())

	app.Get("/", handler.Home)
	app.Get("/search", handler.Search)
	app.Get("/status", handler.Status)
	
	app.Use(middleware.NotFound)
}
