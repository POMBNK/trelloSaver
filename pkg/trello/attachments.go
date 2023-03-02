package trello

import (
	"fmt"
	"net/url"
	"time"
)

type Attachment struct {
	ID        string     `json:"id"`
	Bytes     int64      `json:"bytes"`
	Date      *time.Time `json:"date"`
	EdgeColor string     `json:"edgeСolor"`
	IDMember  string     `json:"idMember"`
	IsUpload  bool       `json:"isUpload"`
	MimeType  string     `json:"mimeType"`
	Name      string     `json:"name"`
	//Preview
	URL      string `json:"url"`
	FileName string `json:"fileName"`
}

// GetAttachmentsFromCard get list of attachments from current Card
func (c *Client) GetAttachmentsFromCard(cardID string) ([]Attachment, error) {
	var attachments []Attachment
	q := url.Values{}
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	path := fmt.Sprintf("cards/%s/attachments", cardID)
	err := c.Get(path, q, &attachments)
	return attachments, err
}
