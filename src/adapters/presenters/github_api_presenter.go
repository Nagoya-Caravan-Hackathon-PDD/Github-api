package presenters

import "github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"

type githubAPIPresenter struct {
}

func NewGithubAPIPresenter() *githubAPIPresenter {
	return &githubAPIPresenter{}
}

func (g *githubAPIPresenter) GithubAPI(err error) (int, output.GithubAPIResponse) {
	if err != nil {
		return 500, output.GithubAPIResponse{}
	}

	return 200, output.GithubAPIResponse{}
}
