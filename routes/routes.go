package routes

import (
	config "BE/configs"
	"BE/database"
	"BE/handler"
	"BE/helper/middleware"
	repoRole "BE/repository/role"
	repoUser "BE/repository/user"

	serviceRole "BE/services/role"
	serviceUser "BE/services/user"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repoUser.NewUserRepository(db)

	svc := serviceUser.NewServiceUser(repo, conf)

	controller := handler.EchoControllerUser{
		Svc: svc,
	}

	e.POST("/register", controller.RegisterHandler)
	e.POST("/login", controller.LoginHandler)

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.GET("/users", controller.GetUsersController)
	adminRoutes.GET("/users/:id", controller.GetUserController)
	adminRoutes.POST("/users/:id", controller.UpdateUserController)
	adminRoutes.DELETE("/users/:id", controller.DeleteUserController)
}

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoRole.NewRoleRepository(db)

	svc := serviceRole.NewServiceRole(repo, conf)

	controller := handler.EchoControllerRole{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/role", controller.CreateRoleController)
	adminRoutes.GET("/role", controller.GetAllRoleController)
	adminRoutes.GET("/role/:id", controller.GetRoleIDController)
	adminRoutes.DELETE("/role/:id", controller.DeleteRoleIDController)
}
