package helper

import (
	config "BE/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTUserMiddleware() echo.MiddlewareFunc {
	secret := config.Config{}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret.JWT_KEY),
		SigningMethod: "HS256",
	})
}
