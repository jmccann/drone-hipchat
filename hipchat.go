package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const notifyURL = "https://api.hipchat.com/v2/room/%s/notification?auth_token=%s"

type Client struct {
	URL string
}

type Message struct {
	From    string `json:"from"`
	Color   string `json:"color"`
	Notify  bool   `json:"notify"`
	Message string `json:"message"`
}

func NewClient(room, token string) *Client {
	return &Client{fmt.Sprintf(notifyURL, room, token)}
}

func (c *Client) Send(msg *Message) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	http.NewRequest("POST", c.URL, buf)
	resp, err := http.Post(c.URL, "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &HipChatError{resp.StatusCode, string(t)}
	}

	return nil
}

type HipChatError struct {
	Code int
	Body string
}

func (e *HipChatError) Error() string {
	return fmt.Sprintf("HipChatError: %d %s", e.Code, e.Body)
}
