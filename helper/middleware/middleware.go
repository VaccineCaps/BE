package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func AuthUser(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, nil
}

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("ini custom masuk", c.Request().Header["Authorization"])

		if err := next(c); err != nil {
			c.Error(err)
		}

		fmt.Println("ini custom keluar")

		return nil
	}
}

func APIKEYMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["X-Api-Key"] == nil {
			return echo.ErrUnauthorized

		}

		apiKey := c.Request().Header["X-Api-Key"][0]
		if apiKey != "ABC" {
			return echo.ErrUnauthorized
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
