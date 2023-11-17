package github_api

import (
	"context"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/cmd/config"
	"github.com/diegosz/go-graphql-client"
	"golang.org/x/oauth2"
)

type Connection struct {
}

func GithubAuthentication() *graphql.Client {
	token := config.Config.Github.GithubToken
	endpoint := config.Config.Github.Endpoint
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient(endpoint, httpClient)
	return client
}
