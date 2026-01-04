package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Token  string
	Client *http.Client
}

func New(token string) *Client {
	return &Client{
		Token:  token,
		Client: &http.Client{},
	}
}

func (c *Client) ApiURL(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", c.Token, method)
}

func (c *Client) Post(method string, payload any, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := c.Client.Post(
		c.ApiURL(method),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}
