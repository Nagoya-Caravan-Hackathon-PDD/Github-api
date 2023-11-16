package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type githubAPIController struct {
	interactors ports.GithubAPIInput
}

func NewGithubAPIController(interactor ports.GithubAPIInput) ports.GithubAPIInput {
	return &githubAPIController{interactor}
}

func (gc *githubAPIController) CreateGitmon(ctx echo.Context) error {
	var reqQuery input.GithubAPIRequest
	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(gc.interactors.CreateGitmon(reqQuery))
}
