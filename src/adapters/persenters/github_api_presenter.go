package persenters

type githubAPIPresenter struct {
}

func NewGithubAPIPresenter() *githubAPIPresenter {
	return &githubAPIPresenter{}
}

// func (g *githubAPIPresenter) GithubAPI(err error) (int, output.GithubAPIResponse) {
// 	if err != nil {
// 		return 500, output.GithubAPIResponse{Message: "NG"}
// 	}
// 	return 200, output.GithubAPIResponse{Message: "OK"}
// }
