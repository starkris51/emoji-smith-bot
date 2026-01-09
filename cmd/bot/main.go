package main

import (
	"emoji-smith-bot/media"
	"emoji-smith-bot/telegram"
	"emoji-smith-bot/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

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

			chatID := update.Message.Chat.ID

			switch {
			case update.Message.Text == "/start":
				_ = client.SendMessage(chatID, "Hello i am the Emoji Smith Bot! Send me an image or video, and I'll convert it into an emoji format that can be uploaded to your emoji pack. \n \nI recommend that images/videos should be the aspect ratio of 1:1, if its 16:9 or 4:3 it will be stretched \n \nVideo emoji will be converted to 3 seconds long make sure your video is less than that, i recommend that you crop before uploading \n \n Made by starkris51")

			case update.Message.Text != "":
				_ = client.SendMessage(chatID, "Please send image or video")

			case update.Message.Document != nil:
				doc := update.Message.Document

				outExt := ""
				var proc func(string, string) error

				mt := strings.ToLower(strings.TrimSpace(doc.MimeType))
				switch {
				case strings.HasPrefix(mt, "image/"):
					outExt = "png"
					proc = media.ProcessImage
				case strings.HasPrefix(mt, "video/"):
					outExt = "webm"
					proc = media.ProcessVideo
				default:
					ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(doc.FileName)), ".")
					switch ext {
					case "png", "jpg", "jpeg", "webp", "gif", "bmp", "tif", "tiff":
						outExt = "png"
						proc = media.ProcessImage
					case "mp4", "mov", "mkv", "webm", "avi", "m4v":
						outExt = "webm"
						proc = media.ProcessVideo
					default:
						_ = client.SendMessage(chatID, "That document doesn't look like an image or video. Please send an image/video file.")
						continue
					}
				}

				if err := utils.HandleTelegramMedia(client, botToken, chatID, doc.FileID, outExt, proc); err != nil {
					log.Println("document handler error:", err)
					_ = client.SendMessage(chatID, err.Error())
				}

			case len(update.Message.Photo) > 0:
				best := update.Message.Photo[len(update.Message.Photo)-1]
				if err := utils.HandleTelegramMedia(client, botToken, chatID, best.FileID, "png", media.ProcessImage); err != nil {
					log.Println("photo handler error:", err)
					_ = client.SendMessage(update.Message.Chat.ID, err.Error())
				}

			case update.Message.Video != nil:
				if err := utils.HandleTelegramMedia(client, botToken, chatID, update.Message.Video.FileID, "webm", media.ProcessVideo); err != nil {
					log.Println("video handler error:", err)
					_ = client.SendMessage(update.Message.Chat.ID, err.Error())
				}

			case update.Message.Animation != nil:
				if err := utils.HandleTelegramMedia(client, botToken, chatID, update.Message.Animation.FileID, "webm", media.ProcessVideo); err != nil {
					log.Println("animation handler error:", err)
					_ = client.SendMessage(update.Message.Chat.ID, err.Error())
				}

			case update.Message.Sticker != nil:
				switch {
				case update.Message.Sticker.IsAnimated && !update.Message.Sticker.IsVideo:
					_ = client.SendMessage(chatID, "Animated .tgs stickers aren't supported yet. Please send an image, video, or a video sticker.")
				case update.Message.Sticker.IsVideo:
					if err := utils.HandleTelegramMedia(client, botToken, chatID, update.Message.Sticker.FileID, "webm", media.ProcessVideo); err != nil {
						log.Println("sticker (video) handler error:", err)
						_ = client.SendMessage(chatID, err.Error())
					}
				default:
					if err := utils.HandleTelegramMedia(client, botToken, chatID, update.Message.Sticker.FileID, "png", media.ProcessImage); err != nil {
						log.Println("sticker (static) handler error:", err)
						_ = client.SendMessage(chatID, err.Error())
					}
				}

			default:
				_ = client.SendMessage(chatID, "Unsupported message type")
			}
		}
	}
}
