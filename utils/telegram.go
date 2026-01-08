package utils

import (
	"emoji-smith-bot/telegram"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type processor func(inputPath, outputPath string) error

func HandleTelegramMedia(client *telegram.Client, botToken string, chatID int64, telegramFileID string, outExt string, process processor) error {
	filePath, err := client.GetFile(telegramFileID)
	if err != nil {
		return fmt.Errorf("getFile: %w", err)
	}

	inExt := strings.TrimPrefix(strings.ToLower(filepath.Ext(filePath)), ".")
	if inExt == "" {
		return fmt.Errorf("could not determine input extension from telegram file_path: %q", filePath)
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, filePath)
	inputPath := fmt.Sprintf("tmp/%s.%s", telegramFileID, inExt)
	outputPath := fmt.Sprintf("output/%s.%s", telegramFileID, outExt)

	if err := DownloadFile(fileURL, inputPath); err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer func() { _ = os.Remove(inputPath) }()

	if err := process(inputPath, outputPath); err != nil {
		return fmt.Errorf("process: %w", err)
	}
	defer func() { _ = os.Remove(outputPath) }()

	if err := client.SendDocument(chatID, outputPath); err != nil {
		return fmt.Errorf("sendDocument: %w", err)
	}

	return nil
}
