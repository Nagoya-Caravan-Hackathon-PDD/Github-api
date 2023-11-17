package controllers

import (
	"fmt"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

// @BasePath /v1
// @Summary Create Gitmon
// @Description Create Gitmon
// @Accept json
// @Produce json
// @Param github_id body string true "Github ID"
// @Success 200 {object} string
// @Router /gitmon [post]
type githubAPIController struct {
	interactor ports.GithubAPIInput
}

func NewGithubAPIController(interactor ports.GithubAPIInput) *githubAPIController {
	return &githubAPIController{interactor: interactor}
}

func (gc *githubAPIController) CreateGitmon(ctx echo.Context) error {
	var reqBody input.GithubAPIRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}
	fmt.Println(reqBody)
	return ctx.JSON(gc.interactor.CreateGitmon(reqBody))
}
