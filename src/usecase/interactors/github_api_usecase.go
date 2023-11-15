package interactors

import (
	"context"

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
	var q struct {
		User struct {
			Name                    string
			Login                   string
			Url                     string
			AvatarUrl               string
			Bio                     string
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int
				}
			}
			Followers struct {
				TotalCount int
			}
		} `graphql:"user(login: $githubID)"`
	}
	if reqQuery.GithubID == "" {
		return i.outputport.GithubAPI(err)
	}
	client := github_api.GithubAuthentication()
	status, err := i.store.GetGithubStatus()
	if err != nil {
		return i.outputport.GithubAPI(err)
	}
	err = client.Query(context.Background(), &q, nil)
}
