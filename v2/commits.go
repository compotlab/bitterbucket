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
	PageLen int             `json:"pagelen"`
	Values  []CommitsValues `json:"values"`
	Next    string          `json:"next"`
}

type CommitsValues struct {
	Hash       string            `json:"hash"`
	Repository CommitsRepository `json:"repository"`
	Links      CommitsLinks      `json:"links"`
	Author     CommitsAuthor     `json:"author"`
	Parents    []CommitsParent   `json:"parents"`
	Date       string            `json:"date"`
	Message    string            `json:"message"`
	Type       string            `json:"type"`
}

type CommitsRepository struct {
	Links    CommitsRepositoryLinks `json:"links"`
	Type     string                 `json:"type"`
	Name     string                 `json:"name"`
	FullName string                 `json:"full_name"`
	UUID     string                 `json:"uuid"`
}

type CommitsRepositoryLinks struct {
	Self   map[string]string `json:"self"`
	HTML   map[string]string `json:"html"`
	Avatar map[string]string `json:"avatar"`
}

type CommitsLinks struct {
	Self       map[string]string `json:"self"`
	Components map[string]string `json:"comments"`
	Patch      map[string]string `json:"patch"`
	HTML       map[string]string `json:"html"`
	Diff       map[string]string `json:"diff"`
	Approve    map[string]string `json:"approve"`
}

type CommitsAuthor struct {
	Raw  string      `json:"raw"`
	User CommitsUser `json:"user"`
}

type CommitsUser struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Type        string    `json:"type"`
	UUID        string    `json:"uuid"`
	Links       UserLinks `json:"links"`
}

type UserLinks struct {
	Self   map[string]string `json:"self"`
	HTML   map[string]string `json:"html"`
	Avatar map[string]string `json:"avatar"`
}

type CommitsParent struct {
	Hash  string      `json:"hash"`
	Type  string      `json:"type"`
	Links ParentLinks `json:"links"`
}

type ParentLinks struct {
	Self map[string]string `json:"self"`
	HTML map[string]string `json:"html"`
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
	_, err := c.Client.Do("GET", u, []byte{}, &data)
	if err != nil {
		return CommitsResponse{}, err
	}

	return data, nil
}
