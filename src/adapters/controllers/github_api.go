package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type githubAPIController struct {
	usecase ports.GithubAPIInput
}

func NewGithubAPIController(usecase ports.GithubAPIInput) *githubAPIController {
	return &githubAPIController{usecase}
}

func (gc *githubAPIController) GetGithubStatus(ctx echo.Context) error {
	var reqQuery input.GithubAPIRequest
	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}
	return nil //ctx.JSON(gc.interactor.GithubAPI(reqQuery))
}
