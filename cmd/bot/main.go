package main

import (
	"emoji-smith-bot/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	botToken := os.Getenv("TOKEN")
	client := telegram.New(botToken)

	offset := 0
	for {
		updates, err := client.GetUpdates(offset)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil {
				continue
			}

			client.SendMessage(update.Message.Chat.ID, "You said: "+update.Message.Text)
		}
	}
}
