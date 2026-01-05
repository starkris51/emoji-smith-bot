package telegram

func (c *Client) SendMessage(chatID int64, text string) error {
	req := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	var resp map[string]any
	return c.Post("sendMessage", req, &resp)
}

func (c *Client) GetFile(fileID string) (string, error) {
	req := map[string]any{
		"file_id": fileID,
	}

	var resp struct {
		filePath string `json:"file_path,omitempty"`
	}
	err := c.Post("getFile", req, &resp)
	return resp.filePath, err
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
