package main

import (
	"emoji-smith-bot/telegram"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("TOKEN")

	if botToken == "" {
		log.Fatalln("missing TOKEN environment variable")
	}

	client := telegram.New(botToken)

	offset := 0

	fmt.Println("Bot is running...")

	for {
		updates, err := client.GetUpdates(offset)
		if err != nil {
			log.Println("telegram error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil {
				continue
			}

			if update.Message.Text != "" {
				client.SendMessage(update.Message.Chat.ID, "Please send image or video")
			} else if len(update.Message.Photo) > 0 {
				for _, photo := range update.Message.Photo {
					client.SendMessage(update.Message.Chat.ID, photo.FileID)
				}
			} else if update.Message.Video.FileID != "" {
				client.SendMessage(update.Message.Chat.ID, "Cool video!")
			} else if update.Message.Animation.FileID != "" {
				client.SendMessage(update.Message.Chat.ID, "Great animation!")
			} else {
				client.SendMessage(update.Message.Chat.ID, "Unsupported message type")
			}
		}
	}
}
