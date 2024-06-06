package handlers

import (
    "github.com/gofiber/fiber/v2"
)

// HelloHandler example
// @Summary Show a hello message
// @Description get string hello
// @ID hello
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /hello [get]
func GetHello(c *fiber.Ctx) error {
    const message = "Hello World!"
    return c.SendString(message)
}