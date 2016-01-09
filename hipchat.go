package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const notifyURL = "https://api.hipchat.com/v2/room/%s/notification?auth_token=%s"

// Client represents the HipChat client.
type Client struct {
	URL string
}

// Message represents the HipChat notification message.
type Message struct {
	From    string `json:"from"`
	Color   string `json:"color"`
	Notify  bool   `json:"notify"`
	Message string `json:"message"`
}

// NewClient returns a new HipChat Client.
func NewClient(room, token string) *Client {
	return &Client{
		URL: fmt.Sprintf(
			notifyURL,
			room,
			token),
	}
}

// Send takes a HipChat notification message and sends it.
func (c *Client) Send(msg *Message) error {
	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	http.NewRequest(
		"POST",
		c.URL,
		buf)

	resp, err := http.Post(
		c.URL,
		"application/json",
		buf)

	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		t, _ := ioutil.ReadAll(resp.Body)
		return &HipChatError{resp.StatusCode, string(t)}
	}

	return nil
}

// HipChatError represents a HipChat error.
type HipChatError struct {
	Code int
	Body string
}

// Error impliments the error interface.
func (e *HipChatError) Error() string {
	return fmt.Sprintf("HipChatError: %d %s", e.Code, e.Body)
}
