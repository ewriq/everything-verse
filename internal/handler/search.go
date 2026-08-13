package handler

import (
	"everything-verse/database"
	

	"github.com/gofiber/fiber/v3"
)

func Search(c fiber.Ctx) error {
	keyword := c.Query("q")

	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query parameter 'q' is required",
		})
	}

	results, err := database.SearchFTS(keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"query":   keyword,
		"results": results,
	})
}
