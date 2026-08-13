package main

import (
	"everything-verse/internal/routes"
	"everything-verse/jobs"
	"time"

	"github.com/gofiber/fiber/v3"

	_ "everything-verse/database"
)

func main() {
	app := fiber.New()

	routes.Api(app)

	go func() {
		counter := time.NewTicker(5 * time.Minute)
		defer counter.Stop()

		jobs.Data() 

		for range counter.C {
			jobs.Data()
		}
	}()

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
