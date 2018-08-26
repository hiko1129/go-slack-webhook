package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Field struct
type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// Attachment struct
type Attachment struct {
	Fallback   *string  `json:"fallback,omitempty"`
	Color      *string  `json:"color,omitempty"`
	PreText    *string  `json:"pretext,omitempty"`
	AuthorName *string  `json:"author_name,omitempty"`
	AuthorLink *string  `json:"author_link,omitempty"`
	AuthorIcon *string  `json:"author_icon,omitempty"`
	Title      *string  `json:"title,omitempty"`
	TitleLink  *string  `json:"title_link,omitempty"`
	Text       *string  `json:"text,omitempty"`
	ImageURL   *string  `json:"image_url,omitempty"`
	ThumbURL   *string  `json:"thumb_url,omitempty"`
	Fields     []*Field `json:"fields,omitempty"`
	Footer     *string  `json:"footer,omitempty"`
	FooterIcon *string  `json:"footer_icon,omitempty"`
	Timestamp  *int64   `json:"ts,omitempty"`
}

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

// AddField func
func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

// Send func
func Send(webhookURL string, payload Payload) (*http.Response, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(b))
	if res != nil {
		res.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
