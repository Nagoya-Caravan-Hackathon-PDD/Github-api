package presenters

import (
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
)

type githubAPIPresenter struct {
}

func NewGithubAPIPresenter() *githubAPIPresenter {
	return &githubAPIPresenter{}
}

func (g *githubAPIPresenter) GithubAPI(args types.CreateGitmon, err error) (int, *output.GithubAPIResponse) {

	if err != nil {
		return 500, nil
	}
	// レスポンス追加
	return http.StatusOK, &output.GithubAPIResponse{
		Owner: args.Owner,
		Name:  args.Name,
	}
}
