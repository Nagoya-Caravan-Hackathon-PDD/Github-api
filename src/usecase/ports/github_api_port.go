package ports

import "github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"

type GithubAPIInput interface {
	CreateGitmon() (int, *output.GithubAPIResponse)
	GetGithubStatus() (string, error)
}

type GithubAPIOutput interface {
	CreateGitmon() (int, *output.GithubAPIResponse)
	GithubAPI(err error) (int, *output.GithubAPIResponse)
}
