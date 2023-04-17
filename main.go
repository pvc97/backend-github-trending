package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/log"
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

	log.Error("Co loi xay ra 2")

	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":1323"))
}
