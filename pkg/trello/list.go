package trello

import (
	"fmt"
	"net/url"
)

type Lists []List

type List struct {
	client  *Client
	ID      string `json:"id"`
	Name    string `json:"name"`
	Closed  bool   `json:"closed"`
	IDBoard string `json:"idBoard,omitempty"`
	Pos     int64  `json:"pos,omitempty"`
	Board   *Board `json:"board,omitempty"`
	Cards   []Card `json:"cards,omitempty"`
}

// GetList method is used to get list from Trello board by ID
func (c *Client) GetList(listID string) (List, error) {
	var list List
	q := url.Values{}
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	path := fmt.Sprintf("lists/%s", listID)
	err := c.Get(path, q, &list)
	if err != nil {
		return list, err
	}
	return list, nil
}

// GetLists method is used to get all created lists from trello board and all meta information about them.
func (c *Client) GetLists(boardID string) (Lists, error) {
	var lists Lists
	q := url.Values{}
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	path := fmt.Sprintf("boards/%s/lists", boardID)
	err := c.Get(path, q, &lists)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
