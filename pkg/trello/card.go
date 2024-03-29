package trello

import (
	"fmt"
	"net/url"
	"time"
)

type Cards []Card

type Card struct {
	client *Client

	//Meta
	ID               string     `json:"id"`
	IDShort          int        `json:"idShort"`
	Name             string     `json:"name"`
	Pos              float32    `json:"pos"`
	Email            string     `json:"email"`
	ShortLink        string     `json:"shortLink"`
	ShortUrl         string     `json:"shortUrl"`
	Url              string     `json:"url"`
	DateLastActivity *time.Time `json:"dateLastActivity"`
	Desc             string     `json:"desc"`
	Due              string     `json:"due"`
	Closed           bool       `json:"closed"`
	Subscribed       bool       `json:"subscribed"`
}

// GetCardsInList Get all cards information in current list
func (c *Client) GetCardsInList(listID string) (Cards, error) {
	var cards []Card
	q := url.Values{}
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	path := fmt.Sprintf("lists/%s/cards", listID)
	err := c.Get(path, q, &cards)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
