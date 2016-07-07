package bitterbucket_test

import (
	"testing"

	"github.com/compotlab/bitterbucket/v2"
)

func TestGetCommits(t *testing.T) {
	c := bitterbucket.Client{
		Auth: bitterbucket.Auth{
			Token: Token,
		},
	}

	req := bitterbucket.CommitsRequest{
		Owner:       Owner,
		Slug:        Slug,
		BranchOrTag: "master",
	}
	_, err := c.Commits().GetCommits(req)
	if err != nil {
		t.Fatal(err)
	}
}
