package server

import (
	config "BE/configs"
	"BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() *echo.Echo {
	app := echo.New()
	app.Use(middleware.CORS())
	conf := config.InitConfiguration()

	routes.RegisterUserGroupAPI(app, conf)
	routes.RegisterRoleGroupAPI(app, conf)
	routes.RegisterProvinceGroupAPI(app, conf)
	routes.RegisterCityGroupAPI(app, conf)
	routes.RegisterHospitalGroupAPI(app, conf)
	routes.RegisterNewsGroupAPI(app, conf)
	routes.RegisterOPsGroupAPI(app, conf)
	routes.RegisterVaccineGroupAPI(app, conf)
	routes.RegisterStokVaccineGroupAPI(app, conf)
	routes.RegisterSessionGroupAPI(app, conf)
	routes.RegisterVaccineTransactionInGroupAPI(app, conf)
	routes.RegisterVaccineTransactionOutGroupAPI(app, conf)
	routes.RegisterAdvertiseGroupAPI(app, conf)
	routes.RegisterCertificateGroupAPI(app, conf)
	return app
}
