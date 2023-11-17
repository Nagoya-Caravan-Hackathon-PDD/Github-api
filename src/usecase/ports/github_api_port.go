package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"
)

type GithubAPIInput interface {
	CreateGitmon(reqQuery input.GithubAPIRequest) (int, *output.GithubAPIResponse)
	//	GetGithubStatus() (string, error)
}

type GithubAPIOutput interface {
	//CreateGitmon(reqQuery input.GithubAPIRequest) (int, *output.GithubAPIResponse)
	GithubAPI(err error) (int, *output.GithubAPIResponse)
}
