package routes

import (
	controllers "dexa-training-crud-karyawan/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func AuthRouter(a *fiber.App) {

	// Create routes group.
	route := a.Group("/token")

	// Routes for GET method:
	route.Get("/new", controllers.NewToken) // create a new access tokens

}
