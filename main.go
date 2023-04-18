package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/log"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"github.com/labstack/echo/v4"
	"os"
)

// init function will be called automatically for each package
func init() {
	err := os.Setenv("APP_NAME", "github")
	if err != nil {
		return
	}

	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "pvc97",
		Password: "1997",
		DbName:   "golang",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
