package trello

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Attachments []Attachment

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
func (c *Client) GetAttachmentsFromCard(cardID string) (Attachments, error) {
	var attachments []Attachment
	q := url.Values{}
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	path := fmt.Sprintf("cards/%s/attachments", cardID)
	err := c.Get(path, q, &attachments)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

// DownloadFile will download an url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (c *Client) DownloadFile(filepath string, url string) error {

	// New GET request with Auth 1.0 headers
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	headerValues := fmt.Sprintf(
		"OAuth oauth_consumer_key=\"%s\","+
			"oauth_token=\"%s\","+
			"oauth_signature_method=\"HMAC-SHA1\","+
			"oauth_version=\"1.0\"", c.Key, c.Token)
	req.Header.Set("Authorization", headerValues)

	// Get the data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	err = os.Mkdir("downloaded", 0700)
	if err != nil && !os.IsExist(err) {
		return err
	}
	out, err := os.Create("downloaded/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
