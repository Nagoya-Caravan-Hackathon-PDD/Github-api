package gateways

import (
	"database/sql"
)

type githubAPIGateway struct {
	db *sql.DB
}

func NewGithubAPIGateway(db *sql.DB) *githubAPIGateway {
	return &githubAPIGateway{db}
}

func (g *githubAPIGateway) GetGithubAPI() error {
	return nil
}
