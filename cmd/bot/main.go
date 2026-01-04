package main

import (
	"emoji-smith-bot/telegram"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	botToken := os.Getenv("TOKEN")
	client := telegram.New(botToken)

	_ = client.Token
}
