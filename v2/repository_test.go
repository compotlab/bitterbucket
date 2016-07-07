package bitterbucket_test

import (
	"testing"

	"github.com/compotlab/bitterbucket/v2"
)

const Owner = ""
const Slug = ""

func TestGetRepository(t *testing.T) {
	c := bitterbucket.Client{
		Auth: bitterbucket.Auth{
			Token: Token,
		},
	}

	req := bitterbucket.RepositoryRequest{
		Owner: Owner,
		Slug:  Slug,
	}
	_, err := c.Repository().GetRepository(req)
	if err != nil {
		t.Fatal(err)
	}
}
