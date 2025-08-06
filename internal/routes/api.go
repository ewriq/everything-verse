package routes

import (
	"everything-verse/internal/handler"

	"github.com/gofiber/fiber/v2"

)

func Api(app fiber.Router) {
	app.Get("/", handler.Home)
	app.Get("/get", handler.Home)
}
