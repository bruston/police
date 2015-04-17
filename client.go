package police

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	USER_AGENT = "v0.1"
	API_URL    = "http://data.police.uk/api/"
)

type Client struct {
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
}

type APIError int

func (e APIError) Error() string {
	return fmt.Sprintf("API returned status code: %d", e)
}

func (c Client) doRequest(method, dst string) (*http.Response, error) {
	req, err := http.NewRequest(method, dst, nil)
	if err != nil {
		return nil, err
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	} else {
		req.Header.Add("User-Agent", USER_AGENT)
	}
	return c.HTTPClient.Do(req)
}

func New() Client {
	return Client{
		BaseURL:    API_URL,
		UserAgent:  USER_AGENT,
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) decodeJSONResponse(method, dst string, target interface{}) error {
	resp, err := c.doRequest(method, c.BaseURL+dst)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return APIError(resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return err
	}
	return nil
}
