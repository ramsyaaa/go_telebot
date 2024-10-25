package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler *Handler) {
	app.Post("/cicd/send-message", handler.HandleMessage)
}
