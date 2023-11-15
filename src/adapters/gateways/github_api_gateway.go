package gateways

import (
	"context"

	"github.com/diegosz/go-graphql-client"
)

type githubAPIGateway struct {
	graph *graphql.Client
}

func NewGithubAPIGateway(graph *graphql.Client) *githubAPIGateway {
	return &githubAPIGateway{graph}
}

type GitHubStatusQuery struct {
	User struct {
		Name                    string
		Login                   string
		Url                     string
		AvatarUrl               string
		Bio                     string
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions int
			}
		}
		Followers struct {
			TotalCount int
		}
	} `graphql:"user(login: $githubID)"`
}

func (g *githubAPIGateway) GetStatus(githubId string) (*GitHubStatusQuery, error) {
	var query GitHubStatusQuery
	if err := g.graph.Query(context.Background(), &query, map[string]interface{}{
		"githubID": graphql.ID(githubId)}); err != nil {
		return nil, err
	}

	return &query, nil
}

func (g *githubAPIGateway) GetGithubStatus(query interface{}) error {
	return nil
}
