package telegram

func (c *Client) SendMessage(chatID int, text string) error {
	req := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	var resp map[string]any
	return c.Post("sendMessage", req, &resp)
}
