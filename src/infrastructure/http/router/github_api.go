package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (router *router) GithubAPI() {
	gc := controllers.NewGithubAPIController(
		interactors.NewGithubAPIInteractor(
			gateways.NewGithubAPIGateway(),
			presenters.NewGithubAPIPresenter(),
		),
	)
	router.echo.GET("/github_api", gc.GithubAPI)
}
