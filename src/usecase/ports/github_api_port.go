package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
)

type GithubAPIInput interface {
	CreateGitmon(reqBody input.GithubAPIRequest) (int, *output.GithubAPIResponse)
	//	GetGithubStatus() (string, error)
}

type GithubAPIOutput interface {
	//CreateGitmon(reqQuery input.GithubAPIRequest) (int, *output.GithubAPIResponse)
	GithubAPI(args types.CreateGitmon, err error) (int, *output.GithubAPIResponse)
}
