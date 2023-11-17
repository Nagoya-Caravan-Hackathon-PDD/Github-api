package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
	"github.com/google/uuid"
)

type GitmonGateway struct {
	db *sql.DB
}

func NewGitmonGateway(db *sql.DB) *GitmonGateway {
	return &GitmonGateway{db: db}
}

func (g *GitmonGateway) Create(arg types.CreateGitmon) error {
	const query = `INSERT INTO gitmons (gitmon_id, owner_id, gitmon_name, exp, base_hp, current_hp, base_attack, current_attack, base_defence, current_defence, base_speed, current_speed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := g.db.Exec(query, uuid.New().String(), arg.Owner, arg.Name, arg.Exp, arg.BaseHP, arg.CurrentHP, arg.BaseAttack, arg.CurrentAttack, arg.BaseDefense, arg.CurrentDefense, arg.BaseSpeed, arg.CurrentSpeed)
	if err != nil {
		return err
	}
	return nil
}
