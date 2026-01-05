package main

import (
	"emoji-smith-bot/telegram"
	"emoji-smith-bot/utils"
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
				downloadID, err := client.GetFile(update.Message.Photo[0].FileID)
				if err != nil {
					log.Println("telegram error:", err)
					continue
				}

				fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, downloadID)
				localPath := fmt.Sprintf("tmp/%s.jpg", update.Message.Photo[0].FileID)

				utils.DownloadFile(fileURL, localPath)
			} else if update.Message.Video.FileID != "" {
			} else if update.Message.Animation.FileID != "" {
				client.SendMessage(update.Message.Chat.ID, "Great animation!")
			} else {
				client.SendMessage(update.Message.Chat.ID, "Unsupported message type")
			}
		}
	}
}
