package ports

import "github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"

type GithubAPIInput interface {
	// GetGithubStatus() (string, error)
}

type GithubAPIOutput interface {
	GithubAPI(err error) (int, output.GithubAPIResponse)
}
