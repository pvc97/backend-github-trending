package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.JWT_KEY),
		// ContextKey: "user", // default
	}

	return middleware.JWTWithConfig(config)
}
