package main

import (
	"emoji-smith-bot/media"
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

	if err := os.MkdirAll("tmp", 0755); err != nil {
		log.Fatalln("failed to create tmp directory:", err)
	}

	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatalln("failed to create output directory:", err)
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
				best := update.Message.Photo[len(update.Message.Photo)-1]

				fileID, err := client.GetFile(best.FileID)
				if err != nil {
					log.Println("telegram error:", err)
					continue
				}

				fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, fileID)
				inputPath := fmt.Sprintf("tmp/%s.jpg", best.FileID)
				outputPath := fmt.Sprintf("output/%s.jpg", best.FileID)

				if err := utils.DownloadFile(fileURL, inputPath); err != nil {
					log.Println("download error:", err)
					continue
				}

				if err := media.ProcessImage(inputPath, outputPath); err != nil {
					log.Println("image processing error:", err)
					continue
				}

				if err := client.SendDocument(update.Message.Chat.ID, outputPath); err != nil {
					log.Println("telegram error:", err)
					continue
				}

				// Delete temporary files
				os.Remove(inputPath)
				os.Remove(outputPath)
			} else if update.Message.Video.FileID != "" {
			} else if update.Message.Animation.FileID != "" {
				client.SendMessage(update.Message.Chat.ID, "Great animation!")
			} else {
				client.SendMessage(update.Message.Chat.ID, "Unsupported message type")
			}
		}
	}
}
