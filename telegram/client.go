package telegram

import (
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

func (c *Client) apiURL(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", c.Token, method)
}
