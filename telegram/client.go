package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
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

func (c *Client) PostMultipart(method string, fields map[string]string, fileField, fileName string, fileBytes []byte, out any) error {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	for k, v := range fields {
		if err := w.WriteField(k, v); err != nil {
			_ = w.Close()
			return err
		}
	}

	fw, err := w.CreateFormFile(fileField, fileName)
	if err != nil {
		_ = w.Close()
		return err
	}
	if _, err := fw.Write(fileBytes); err != nil {
		_ = w.Close()
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.ApiURL(method), &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}
