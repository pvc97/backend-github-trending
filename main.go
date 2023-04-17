package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "1997",
		DbName: "golang",
	}

	sql.Connect()

	defer sql.Close()

	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":1323"))
}
