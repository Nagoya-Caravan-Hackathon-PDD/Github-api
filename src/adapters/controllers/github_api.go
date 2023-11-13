package controllers

import (
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type githubAPIController struct {
	usecase ports.GithubAPIInput
}

func NewGithubAPIController(usecase ports.GithubAPIInput) *githubAPIController {
	return &githubAPIController{usecase}
}

func (g *githubAPIController) GithubAPI(ctx echo.Context) error {
	var reqQuery input.GithubAPIRequest

	if err := ctx.Bind(&reqQuery); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(g.usecase.GithubAPI(reqQuery))
}
