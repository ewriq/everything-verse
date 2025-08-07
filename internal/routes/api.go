package routes

import (
	"everything-verse/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Api(app fiber.Router) {
	app.Get("/", handler.Home)
	app.Get("/search", handler.Search)
	app.Get("/collect", handler.Collect)
	app.Get("/status", handler.Status)
}
