package police

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	USER_AGENT = "v0.1"
	API_URL    = "http://data.police.uk/api/"
)

type Client struct {
	baseURL    string
	UserAgent  string
	HTTPClient http.Client
}

func (c *Client) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	} else {
		req.Header.Add("User-Agent", USER_AGENT)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func New() Client {
	return Client{
		baseURL:    API_URL,
		UserAgent:  USER_AGENT,
		HTTPClient: http.Client{Timeout: time.Second * 10},
	}
}

func (c Client) decodeJSONResponse(dst string, target interface{}) error {
	resp, err := c.get(c.baseURL + dst)
	if err != nil {
		return fmt.Errorf("unable to connect to API: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return fmt.Errorf("unable to unmarshal JSON response: %s", err)
	}
	return nil
}
