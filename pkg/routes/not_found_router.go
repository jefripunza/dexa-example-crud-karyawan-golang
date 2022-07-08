package routes

import (
	"github.com/gofiber/fiber/v2"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			method := c.Method()

			if method != "GET" && method != "POST" && method != "PATCH" && method != "DELETE" {
				// Return HTTP 404 status and JSON response.
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "forbidden",
				})
			}

			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "endpoint not found!",
			})
		},
	)
}
