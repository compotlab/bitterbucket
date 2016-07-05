package bitterbucket

import "fmt"

type Repository struct {
	Client *Client
}

type RepositoryRequest struct {
	Owner string
	Slug  string
}

type RepositoryResponse struct {
	Links RepositoryResponseLinks `json:"links"`
	Name  string                  `json:"name"`
	Size  int                     `json:"size"`
}

type RepositoryResponseLinks struct {
	Clone []map[string]string
}

func (c *Client) Repository() *Repository {
	return &Repository{Client: c}
}

func (c *Repository) GetRepository(r RepositoryRequest) (RepositoryResponse, error) {
	u := fmt.Sprintf("%s/repositories/%s/%s?access_token=%s", APIURL, r.Owner, r.Slug, c.Client.Auth.Token)

	data := RepositoryResponse{}
	_, err := c.Client.APICall(u, &data)
	if err != nil {
		return RepositoryResponse{}, err
	}

	return data, nil
}
