package routes

import (
	controllers "dexa-training-crud-karyawan/pkg/controllers"
	middleware "dexa-training-crud-karyawan/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func KaryawanManagementRouter(a *fiber.App) {

	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for method:
	route.Post("/karyawan", middleware.JWTProtected(), controllers.CreateKaryawan)                                       // menambah karyawan baru
	route.Get("/karyawan/:show/:page/:order_by/:keyword", middleware.JWTProtected(), controllers.ShowKaryawanPagination) // list data karyawan (pagination)
	route.Patch("/karyawan", middleware.JWTProtected(), controllers.UpdateKaryawan)                                      // update karyawan
	route.Delete("/karyawan", middleware.JWTProtected(), controllers.DeleteKaryawan)                                     // delete karyawan

}
