package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Payload struct
type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

// Client struct
type Client struct {
	webhookURL string
}

// New func
func New(webhookURL string) *Client {
	return &Client{webhookURL: webhookURL}
}

// Send func
func (c *Client) Send(payload Payload) (*http.Response, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(c.webhookURL, "application/json", bytes.NewBuffer(b))
	if res != nil {
		res.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
