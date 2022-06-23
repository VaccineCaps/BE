package routes

import (
	config "BE/configs"
	"BE/database"
	"BE/helper/middleware"

	// Deklarasi Handler 
	handlerUser 	"BE/handler/user"
	handlerRole 	"BE/handler/role"
	handlerProvince "BE/handler/province"
	handlerCities	"BE/handler/cities"
	handlerHospital "BE/handler/hospital"
	handlerNews		"BE/handler/news"
	handlerOP 		"BE/handler/otherperson"
	handlerVStatus	"BE/handler/vaccine_status"

	// Deklarasi repository 
	repoRole 		"BE/repository/role"
	repoUser 		"BE/repository/user"
	repoProvince 	"BE/repository/province"
	repoCity 		"BE/repository/cities"
	repoHospital 	"BE/repository/hospital"
	repoNews 		"BE/repository/news"
	repoOP			"BE/repository/otherperson"
	repoVStatus 	"BE/repository/vaccine_status"

	// Deklarasi Services 
	serviceRole 		"BE/services/role"
	serviceUser 		"BE/services/user"
	serviceProvince 	"BE/services/province"
	serviceCity 		"BE/services/cities"
	serviceHospital 	"BE/services/hospitals"
	serviceNews			"BE/services/news"
	serviceOP 			"BE/services/otherperson"
	serviceVStatus			"BE/services/vaccine_status"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repoUser.NewUserRepository(db)

	svc := serviceUser.NewServiceUser(repo, conf)

	controller := handlerUser.EchoControllerUser{
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

	controller := handlerRole.EchoControllerRole{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/role", controller.CreateRoleController)
	adminRoutes.GET("/role", controller.GetAllRoleController)
	adminRoutes.GET("/role/:id", controller.GetRoleIDController)
	adminRoutes.DELETE("/role/:id", controller.DeleteRoleIDController)
}

func RegisterProvinceGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoProvince.NewProvinceRepository(db)

	svc := serviceProvince.NewServiceProvince(repo, conf)

	controller := handlerProvince.EchoControllerProvince{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/province", controller.CreateProvinceController)
	adminRoutes.GET("/province", controller.GetAllProvinceController)
	adminRoutes.GET("/province/:id", controller.GetProvinceIDController)
	adminRoutes.DELETE("/province/:id", controller.DeleteProvinceIDController)
}

func RegisterCityGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoCity.NewCityRepository(db)

	svc := serviceCity.NewServiceCity(repo, conf)

	controller := handlerCities.EchoControllerCity{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/cities", controller.CreateCityController)
	adminRoutes.GET("/cities", controller.GetAllCityController)
	adminRoutes.GET("/cities/:id", controller.GetCityIDController)
	adminRoutes.DELETE("/cities/:id", controller.DeleteCityIDController)
}

func RegisterHospitalGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoHospital.NewHospitalRepository(db)

	svc := serviceHospital.NewServiceHospitals(repo, conf)

	controller := handlerHospital.EchoControllerHospital{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/hospitals", controller.CreateHospitalController)
	adminRoutes.GET("/hospitals", controller.GetHospitalController)
	adminRoutes.GET("/hospitals/:id", controller.GetHospitalIDController)
	adminRoutes.PUT("/hospitals/:id", controller.UpdateHospitalController)
	adminRoutes.DELETE("/hospitals/:id", controller.DeleteHospitalController)

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.GET("/hospitals", controller.GetHospitalController)
	userRoutes.GET("/hospitals/:id", controller.GetHospitalIDController)

}

func RegisterNewsGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoNews.NewNewsRepository(db)

	svc := serviceNews.NewServiceNews(repo, conf)

	controller := handlerNews.EchoControllerNews{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/news", controller.CreateNewsController)
	adminRoutes.GET("/news", controller.GetNewsController)
	adminRoutes.GET("/news/:id", controller.GetNewsIDController)
	adminRoutes.PUT("/news/:id", controller.UpdateNewsController)
	adminRoutes.DELETE("/news/:id", controller.DeleteNewsController)
}

func RegisteOPsGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoOP.NewOtherRepository(db)

	svc := serviceOP.NewServiceOtherPerson(repo, conf)

	controller := handlerOP.EchoControllerOther{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/others", controller.CreateOtherController)
	adminRoutes.GET("/others", controller.GetOtherController)
	adminRoutes.GET("/others/:id", controller.GetOtherIDController)
	adminRoutes.PUT("/others/:id", controller.UpdateOtherController)
	adminRoutes.DELETE("/others/:id", controller.DeleteOtherController)

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.POST("/others", controller.CreateOtherController)
	userRoutes.GET("/others", controller.GetOtherController)
	userRoutes.GET("/others/:id", controller.GetOtherIDController)
}

func RegisterVStatussGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoVStatus.NewVStatusRepository(db)

	svc := serviceVStatus.NewServiceVStatus(repo, conf)

	controller := handlerVStatus.EchoControllerVStatus{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/vstatus", controller.CreateVStatusController)
	adminRoutes.GET("/vstatus", controller.GetAllVStatusController)
	adminRoutes.GET("/vstatus/:id", controller.GetVStatusIDController)
	adminRoutes.PUT("/vstatus/:id", controller.UpdateVStatusController)
	adminRoutes.DELETE("/vstatus/:id", controller.DeleteVStatusIDController)

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.GET("/vstatus", controller.GetAllVStatusController)
	userRoutes.GET("/vstatus/:id", controller.GetVStatusIDController)
}