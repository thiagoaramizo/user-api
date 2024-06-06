package middleware

import (
    "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // lógica de autenticação
    return c.Next()
}