package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Mike",
		"email": "cuongpv2209@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json:"email"` // convert Email to email
		FullName string `json:"fullName"`
	}

	user := User{
		Email:    "cuongpv2209@gmail.com",
		FullName: "Pham Cuong",
	}
	return c.JSON(http.StatusOK, user)
}
