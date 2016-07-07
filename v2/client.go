package bitterbucket

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// APIURL bitbucket api url
const APIURL = "https://api.bitbucket.org/2.0"

// ErrHTTPUnauthorized unauthorized http error
var ErrHTTPUnauthorized = errors.New("http: unauthorized")

// Client api client
type Client struct {
	Auth Auth
}

// Auth api client authorization data
type Auth struct {
	Token string
}

// Do make http request, unmarshal response data and return http response
func (c *Client) Do(method string, u string, b []byte, data interface{}) (*http.Response, error) {
	buf := bytes.NewReader(b)

	r, err := http.NewRequest("GET", u, buf)
	if err != nil {
		return &http.Response{}, err
	}

	if len(b) > 0 {
		r.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(r)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &http.Response{}, err
	}

	if resp.StatusCode == 401 {
		return &http.Response{}, ErrHTTPUnauthorized
	}

	err = json.Unmarshal(body, data)
	return resp, err
}
