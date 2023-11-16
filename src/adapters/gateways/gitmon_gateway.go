package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
)

type GitmonGateway struct {
	db *sql.DB
}

func NewGitmonGateway(db *sql.DB) *GitmonGateway {
	return &GitmonGateway{db: db}
}

func (g *GitmonGateway) Create(arg types.CreateGitmon) error {

}
