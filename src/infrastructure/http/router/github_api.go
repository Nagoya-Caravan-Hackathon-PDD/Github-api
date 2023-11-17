package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/interactors"
)

func (router *router) GithubAPI() {
	g := router.echo.Group("/v1")
	gc := controllers.NewGithubAPIController(
		interactors.NewGithubAPIInteractor(
			gateways.NewGithubAPIGateway(router.graphqlClient),
			presenters.NewGithubAPIPresenter(),
		),
	)
	router.echo.POST("/gitmon", gc.GetGithubStatus)
}
