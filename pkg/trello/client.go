package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const basedURL = "api.trello.com"

// TODO LIST:
//		Todo: fix hardcoded constant params:key,token in all client's methods.

type Client struct {
	Client  *http.Client
	Token   string
	Key     string
	BoardID string
	BaseURL string
}

// New Trello client constructor
func New(token, key, boardID string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		Token:   token,
		Key:     key,
		BoardID: boardID,
		BaseURL: basedURL,
	}
}

// {{BasedURL}}boards/{{BoardID}}/lists?key={{Key}}&token={{Token}}
// <basedURL> <----------path----------> <--------------params----->
// Path example: "/members/me/boards"

// Get used to make a GET http request in all trello client's methods and get response that
// unmarshalled from json to trello objects.
func (c *Client) Get(path string, q url.Values, respBody interface{}) error {
	u := url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "1/" + path,
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("Bad GET request %s", err)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("Can't DO request %s", err)
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP Read error on response: %s", err)
	}

	err = json.Unmarshal(b, respBody)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
