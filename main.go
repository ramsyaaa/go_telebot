package main

import (
	"go_telebot/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables instead")
	}

	// config.Connect()
	config.Route()
}
