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
	const query = `INSERT INTO gitmons (gitmon_name, exp, base_hp, base_attack, base_defense, base_speed) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := g.db.Exec(query, arg.Name, arg.Exp, arg.HP, arg.Attack, arg.Defense, arg.Speed)
	if err != nil {
		return err
	}
	return nil
}
