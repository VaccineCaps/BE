package routes

import (
	config "BE/configs"
	"BE/database"
	"BE/handler"
	"BE/helper/middleware"
	repository "BE/repository/user"
	service "BE/services/user"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewUserRepository(db)

	svc := service.NewServiceUser(repo, conf)

	controller := handler.EchoControllerUser{
		Svc: svc,
	}

	e.POST("/register", controller.RegisterHandler)
	e.POST("/login", controller.LoginHandler)

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.GET("/users", controller.GetUsersController)
}
