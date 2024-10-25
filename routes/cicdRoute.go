package routes

import (
	"go_telebot/modules/cicd/http"

	"github.com/gofiber/fiber/v2"
)

func CICDRouter(app *fiber.App) {
	http.SetupRoutes(app, http.NewHandler())
}
