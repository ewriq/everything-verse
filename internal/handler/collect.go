package handler

import (
	"everything-verse/database"
	"everything-verse/jobs"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Collect(c *fiber.Ctx) error {
	count, err := database.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
	}


	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("PANIC in manual data collection: %v\n", r)
			}
		}()

		fmt.Println("🔄 Manual data collection triggered")
		jobs.Data()
	}()

	return c.JSON(fiber.Map{
		"status":        "OK",
		"message":       "Data collection started",
		"current_count": len(count),
	})
}
