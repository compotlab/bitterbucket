package bitterbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const APIURL = "https://api.bitbucket.org/2.0"

type Client struct {
	Auth Auth
}

type Auth struct {
	Token string
}

func (c *Client) APICall(u string, data interface{}) (*http.Response, error) {
	r, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return &http.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, data)
	return resp, err
}
