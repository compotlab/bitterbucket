package bitterbucket

import (
	"fmt"
)

type Commits struct {
	Client *Client
}

type CommitsRequest struct {
	Owner       string
	Slug        string
	BranchOrTag string
	Include     []string
	Exclude     []string
}

type CommitsResponse struct {
	PageLen int                     `json:"pagelen"`
	Values  []CommitsResponseValues `json:"values"`
	Next    string                  `json:"next"`
}

type CommitsResponseValues struct {
	Hash       string                          `json:"hash"`
	Repository CommitsResponseValuesRepository `json:"repository"`
	//Links      string `json:"links"`
	//Author     string `json:"author"`
	//Parents    string `json:"parents"`
	Date    string `json:"date"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type CommitsResponseValuesRepository struct {
	//Links string        `json:"links"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	UUID     string `json:"uuid"`
}

func (c *Client) Commits() *Commits {
	return &Commits{Client: c}
}

func (c *Commits) GetCommits(cr CommitsRequest) (CommitsResponse, error) {
	u := fmt.Sprintf("%s/repositories/%s/%s/commits", APIURL, cr.Owner, cr.Slug)

	if cr.BranchOrTag != "" {
		u = fmt.Sprintf("%s/%s", u, cr.BranchOrTag)
	}

	if c.Client.Auth.Token != "" {
		u = fmt.Sprintf("%s?access_token=%s", u, c.Client.Auth.Token)
	}

	return c.getCommits(u)
}

func (c *Commits) GetNextCommits(cr CommitsResponse) (CommitsResponse, error) {
	if cr.Next == "" {
		return CommitsResponse{}, nil
	}

	return c.getCommits(cr.Next)
}

func (c *Commits) getCommits(u string) (CommitsResponse, error) {
	data := CommitsResponse{}
	_, err := c.Client.APICall(u, &data)
	if err != nil {
		return CommitsResponse{}, err
	}

	return data, nil
}
