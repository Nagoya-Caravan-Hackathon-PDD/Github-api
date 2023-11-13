package interactors

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/github_api"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
)

type GithubAPIInteractor struct {
	store      dai.GithubAPIDai
	outputport ports.GithubAPIOutput
}

func NewGithubAPIInteractor(store dai.GithubAPIDai, outputport ports.GithubAPIOutput) *GithubAPIInteractor {
	return &GithubAPIInteractor{store, outputport}
}

func (i *GithubAPIInteractor) GetGithubStatus(reqQuery input.GithubAPIRequest) (int, output.GithubAPIResponse) {
	var err error
	client := github_api.GithubAuthentication()
	if reqQuery.GithubID != "" {
		status, err := i.store.GetGithubStatus()
	}

}
