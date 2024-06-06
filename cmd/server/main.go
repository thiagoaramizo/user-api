package main

import (
	_ "gerenciamento-servicos/docs" // Import docs package for Swagger
	"gerenciamento-servicos/internal/config"
	"gerenciamento-servicos/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for a Fiber application.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
    app := fiber.New()

    // Route to serve the Swagger documentation
    app.Get("/swagger/*", swagger.HandlerDefault)

    config.LoadConfig()
    routes.SetupRoutes(app)

    app.Listen(":3000")
}