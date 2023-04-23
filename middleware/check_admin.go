package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := req.SignIn{}

			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}

			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Bạn không không có quyền gọi api này !",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}
