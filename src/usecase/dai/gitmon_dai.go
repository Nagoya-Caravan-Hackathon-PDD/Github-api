package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"

type GitmonDai interface {
	// Createのインターフェースを定義する
	Create(types.CreateGitmon) error
}
