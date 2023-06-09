package router

import (
	"backend-github-trending/handler"
	middleware "backend-github-trending/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	// profile
	user := api.Echo.Group("/user", middleware.JwtMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile", api.UserHandler.UpdateProfile)

	// GitHub repo
	github := api.Echo.Group("/github", middleware.JwtMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)

	// bookmark
	bookmark := api.Echo.Group("/bookmark", middleware.JwtMiddleware())
	bookmark.GET("/list", api.RepoHandler.SelectBookmarks)
	bookmark.POST("/add", api.RepoHandler.Bookmark)
	bookmark.DELETE("/delete", api.RepoHandler.DelBookmark)
}
