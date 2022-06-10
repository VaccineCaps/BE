package server

import (
	config "BE/configs"
	"BE/routes"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := config.InitConfiguration()

	routes.RegisterUserGroupAPI(app, conf)
	routes.RegisterRoleGroupAPI(app, conf)
	return app
}
