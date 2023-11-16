package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"

type GithubAPIDai interface {
	GetStatus(githubId string) (*types.GitHubStatusQuery, error)
}
