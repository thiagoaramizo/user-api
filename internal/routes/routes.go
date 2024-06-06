package routes

import (
    "github.com/gofiber/fiber/v2"
    "gerenciamento-servicos/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHello)
    app.Get("/user/:id", handlers.GetUser)
}