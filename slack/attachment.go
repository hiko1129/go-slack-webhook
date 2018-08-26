package slack

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

// AddField func
func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}
