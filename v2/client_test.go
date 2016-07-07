package bitterbucket_test

import (
	"net/http"
	"testing"

	"github.com/golang/oauth2/bitbucket"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const Token = ""

func TestAuthorization(t *testing.T) {
	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     bitbucket.Endpoint,
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{})
	_, err := conf.Exchange(ctx, "code")
	if err != nil {
		t.Fatal(err)
	}
}
