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
	SCM         string            `json:"scm"`
	Website     string            `json:"website"`
	HasWiki     bool              `json:"has_wiki"`
	Name        string            `json:"name"`
	Links       RepositoryLinks   `json:"links"`
	ForkPolicy  string            `json:"fork_policy"`
	UUID        string            `json:"uuid"`
	Project     RepositoryProject `json:"project"`
	Language    string            `json:"language"`
	CreatedOn   string            `json:"created_on"`
	FullName    string            `json:"full_name"`
	HasIssues   bool              `json:"has_issues"`
	Owner       RepositoryOwner   `json:"owner"`
	UpdatedOn   string            `json:"updated_on"`
	Size        int               `json:"size"`
	Type        string            `json:"type"`
	IsPrivate   bool              `json:"is_private"`
	Description string            `json:"description"`
}

type RepositoryLinks struct {
	Watchers     map[string]string   `json:"watchers"`
	Branches     map[string]string   `json:"branches"`
	Tags         map[string]string   `json:"tags"`
	Commits      map[string]string   `json:"commits"`
	Clone        []map[string]string `json:"clone"`
	Self         map[string]string   `json:"self"`
	HTML         map[string]string   `json:"html"`
	Avatar       map[string]string   `json:"avatar"`
	Hooks        map[string]string   `json:"hooks"`
	Forks        map[string]string   `json:"forks"`
	Downloads    map[string]string   `json:"downloads"`
	PullRequests map[string]string   `json:"pullrequests"`
}

type RepositoryProject struct {
	Key   string       `json:"key"`
	Type  string       `json:"type"`
	UUID  string       `json:"uuid"`
	Links ProjectLinks `json:"links"`
	Name  string       `json:"name"`
}

type ProjectLinks struct {
	HTML   map[string]string `json:"html"`
	Avatar map[string]string `json:"avatar"`
}

type RepositoryOwner struct {
	Username    string     `json:"username"`
	DisplayName string     `json:"display_name"`
	Type        string     `json:"type"`
	UUID        string     `json:"uuid"`
	Links       OwnerLinks `json:"links"`
}

type OwnerLinks struct {
	Self   map[string]string `json:"self"`
	HTML   map[string]string `json:"html"`
	Avatar map[string]string `json:"avatar"`
}

func (c *Client) Repository() *Repository {
	return &Repository{Client: c}
}

func (c *Repository) GetRepository(r RepositoryRequest) (RepositoryResponse, error) {
	u := fmt.Sprintf("%s/repositories/%s/%s?access_token=%s", APIURL, r.Owner, r.Slug, c.Client.Auth.Token)

	data := RepositoryResponse{}
	_, err := c.Client.Do("GET", u, []byte{}, &data)
	if err != nil {
		return RepositoryResponse{}, err
	}

	return data, nil
}
