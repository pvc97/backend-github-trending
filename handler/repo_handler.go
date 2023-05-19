package handler

import (
	"backend-github-trending/model"
	"backend-github-trending/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RepoHandler struct {
	GithubRepo repository.GithubRepo
}

func (r RepoHandler) RepoTrending(c echo.Context) error {
	repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), 25)
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       repos,
	})
}
