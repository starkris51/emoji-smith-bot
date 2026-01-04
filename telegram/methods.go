package telegram

func (c *Client) SendMessage(chatID int64, text string) error {
	req := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	var resp map[string]any
	return c.Post("sendMessage", req, &resp)
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
