package hasura

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestBody struct {
	Type    string      `json:"type"`
	Version uint        `json:"version,omitempty"`
	Args    interface{} `json:"args"`
}

type Client struct {
	Endpoint    string
	AdminSecret string
	Client      *http.Client
}

func (r *RequestBody) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (c *Client) Request(body []byte) (*http.Response, error) {

	req, err := http.NewRequest(
		http.MethodPost,
		c.Endpoint+"/v1/query",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Hasura-Admin-Secret", c.AdminSecret)

	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.Client.Do(req)
}