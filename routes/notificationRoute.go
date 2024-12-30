package routes

import (
	"go_telebot/modules/notification/http"

	"github.com/gofiber/fiber/v2"
)

func NotificationRouter(app *fiber.App) {
	http.SetupRoutes(app, http.NewHandler())
}
