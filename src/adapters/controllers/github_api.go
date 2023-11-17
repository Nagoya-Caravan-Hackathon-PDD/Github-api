package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type githubAPIController struct {
	interactor ports.GithubAPIInput
}

func NewGithubAPIController(interactor ports.GithubAPIInput) *githubAPIController {
	return &githubAPIController{interactor: interactor}
}

// Gitmon godoc
// @Summary		Create Gitmon
// @Description	Create Gitmon
// @Accept			json
// @Produce		json
// @Param			GithubAPIRequest		body		input.GithubAPIRequest	true	"Github ID"
// @Success		200				{object}	output.GithubAPIResponse
// @Failure		500				{object}	string
// @Router			/v1/gitmon [post]
func (gc *githubAPIController) CreateGitmon(ctx echo.Context) error {
	var reqBody input.GithubAPIRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(gc.interactor.CreateGitmon(reqBody))
}
