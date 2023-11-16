package gateways

import (
	"context"
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
	"github.com/diegosz/go-graphql-client"
)

type githubAPIGateway struct {
	graph *graphql.Client
}

func NewGithubAPIGateway(graph *graphql.Client) *githubAPIGateway {
	return &githubAPIGateway{graph}
}

func (g *githubAPIGateway) GetStatus(githubId string) (*types.GitHubStatusQuery, error) {
	var query types.GitHubStatusQuery
	if err := g.graph.Query(context.Background(), &query, map[string]interface{}{
		"githubID": graphql.String(githubId)}); err != nil {
		return nil, err
	}
	log.Printf("query: %+v", query)
	return &query, nil
}

func (g *githubAPIGateway) GetGithubStatus(query interface{}) error {
	return nil
}
