package helper

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckTokenAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header["Authorization"]
		if len(authHeader) < 1 {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "token required",
				"debug_message": "token required",
			})
		}
		token := strings.Split(authHeader[0], " ")[1]

		userClaim, err := ClaimJwt(token)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "token not valid",
				"debug_message": err.Error(),
			})
		}

		// check admin
		if userClaim.RoleId != 1 {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "unauthorization",
				"debug_message": "unauthorization",
			})
		}

		return next(c)
	}
}

func CheckTokenUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header["Authorization"]
		if len(authHeader) < 1 {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "token required",
				"debug_message": "token required",
			})
		}
		token := strings.Split(authHeader[0], " ")[1]

		userClaim, err := ClaimJwt(token)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "token not valid",
				"debug_message": err.Error(),
			})
		}

		// check admin
		if userClaim.RoleId != 2 {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"error":         true,
				"message":       "unauthorization",
				"debug_message": "unauthorization",
			})
		}

		return next(c)
	}
}
