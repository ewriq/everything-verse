package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
)


var Compress = compress.New(compress.Config{
	Level: compress.LevelBestSpeed,
})

var Cors = cors.New(cors.Config{
	AllowOrigins: []string{"https://gofiber.io", "https://gofiber.net"},
	AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
})

func Error(c fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": "500: Internal server error",
	})
}

var Logger = logger.New(logger.Config{
	Format:     "${time} | ${pid} | ${latency} | ${status} - ${method} ${path} | ${ip}\n",
	TimeFormat: "02.01.2006 15:04:05",
})

func NotFound(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"code":    fiber.StatusNotFound,
		"message": "Not Found",
	})
}

var RateLimit = limiter.New(limiter.Config{
	Max:        1000,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit detected.",
		})
	},
})

func Security(c fiber.Ctx) error {
	c.Set("Content-Security-Policy", "default-src * 'self' data: blob: 'unsafe-inline' 'unsafe-eval'")
	return c.Next()
}
