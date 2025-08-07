package handler

import (
	"everything-verse/database"

	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	
	data, err := database.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Database error",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":        "OK",
		"total_records": len(data),
		"database":      "SQLite",
	})
}
