package main

import (
	"dexa-training-crud-karyawan/pkg/configs"

	middleware "dexa-training-crud-karyawan/pkg/middlewares"
	routes "dexa-training-crud-karyawan/pkg/routes"
	run "dexa-training-crud-karyawan/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// nggak mau tau pokoknya paling awal :'D wkwkwkkk
	godotenv.Load(".env")

	// Define Fiber config.
	config := configs.FiberConfig()
	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// import middleware yang dibutuhkan
	middleware.Cors(app)
	middleware.Logger(app)

	// import semua router yang dibutuhkan
	routes.SwaggerRoute(app)
	routes.AuthRouter(app)
	routes.KaryawanManagementRouter(app)
	routes.NotFoundRoute(app)

	run.DatabaseTest()
	run.Server(app) // nggak mau tau pokoknya paling belakang :')

}
