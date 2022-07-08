package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware

func Cors(a *fiber.App) {
	// Add CORS to each route.
	a.Use(
		cors.New(),
	)
}

func Logger(a *fiber.App) {
	// Add simple logger.
	a.Use(
		logger.New(),
	)
}
