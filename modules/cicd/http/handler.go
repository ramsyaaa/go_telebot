package http

import (
	"log"
	"os"

	"go_telebot/helper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Handler struct {
	Bot *tgbotapi.BotAPI
}

func NewHandler() *Handler {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Bot token is not set. Please set the TELEGRAM_BOT_TOKEN environment variable.")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Handler{
		Bot: bot,
	}
}

func (h *Handler) HandleMessage(c *fiber.Ctx) error {
	var requestBody map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(helper.APIResponse("Invalid request", 400, "error", nil))
	}

	message, ok := requestBody["message"]
	if !ok {
		return c.Status(400).JSON(helper.APIResponse("Message not found in request body", 400, "error", nil))
	}

	msg := tgbotapi.NewMessageToChannel("@cicd_flightapp", message)
	_, err := h.Bot.Send(msg)
	if err != nil {
		return c.Status(500).JSON(helper.APIResponse("Gagal mengirim pesan", 500, "error", nil))
	}

	return c.JSON(helper.APIResponse("Pesan berhasil dikirim", 200, "success", message))
}
