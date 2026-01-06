package telegram

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Client) SendMessage(chatID int64, text string) error {
	req := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	var resp map[string]any
	return c.Post("sendMessage", req, &resp)
}

func (c *Client) SendDocument(chatID int64, filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fields := map[string]string{
		"chat_id": fmt.Sprint(chatID),
	}

	var resp map[string]any
	return c.PostMultipart("sendDocument", fields, "document", filepath.Base(filePath), b, &resp)

}

func (c *Client) GetFile(fileID string) (string, error) {
	req := map[string]any{
		"file_id": fileID,
	}

	var resp struct {
		Ok     bool `json:"ok"`
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}

	err := c.Post("getFile", req, &resp)
	return resp.Result.FilePath, err
}

func (c *Client) GetUpdates(offset int) ([]Update, error) {
	req := map[string]any{
		"offset":  offset,
		"timeout": 30,
	}

	var resp struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	err := c.Post("getUpdates", req, &resp)
	return resp.Result, err
}
