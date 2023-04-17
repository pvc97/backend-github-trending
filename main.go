package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"runtime"
)

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

	logError("Co loi xay ra")

	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":1323"))
}

func logError(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(errMsg)
}
