package routes

import (
	config "BE/configs"
	"BE/database"
	"BE/helper/middleware"

	// Deklarasi Handler
	handlerCities "BE/handler/cities"
	handlerHospital "BE/handler/hospital"
	handlerNews "BE/handler/news"
	handlerOP "BE/handler/otherperson"
	handlerProvince "BE/handler/province"
	handlerRole "BE/handler/role"
	handlerUser "BE/handler/user"
	handlerVaccine "BE/handler/vaccine"
	handlerVaccineStok "BE/handler/vaccinehospital"

	// Deklarasi repository
	repoCity "BE/repository/cities"
	repoHospital "BE/repository/hospital"
	repoNews "BE/repository/news"
	repoOP "BE/repository/otherperson"
	repoProvince "BE/repository/province"
	repoRole "BE/repository/role"
	repoUser "BE/repository/user"
	repoVaccine "BE/repository/vaccine"
	repoStokVaccine "BE/repository/vaccinehospital"

	// Deklarasi Services
	serviceCity "BE/services/cities"
	serviceHospital "BE/services/hospitals"
	serviceNews "BE/services/news"
	serviceOP "BE/services/otherperson"
	serviceProvince "BE/services/province"
	serviceRole "BE/services/role"
	serviceUser "BE/services/user"
	serviceVaccine "BE/services/vaccine"
	serviceStokVaccine "BE/services/vaccinehospital"

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
	adminRoutes.GET("/hospitals", controller.GetHospitalController)
	adminRoutes.GET("/hospitals/:id", controller.GetHospitalIDController)

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

func RegisterOPsGroupAPI(e *echo.Echo, conf config.Config) {
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
}

func RegisterVaccineGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoVaccine.NewVaccineRepository(db)

	svc := serviceVaccine.NewServiceVaccine(repo, conf)

	controller := handlerVaccine.EchoControllerVaccine{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/vaccine", controller.CreateVaccineController)
	adminRoutes.GET("/vaccine", controller.GetAllVaccineController)
	adminRoutes.GET("/vaccine/:id", controller.GetVaccineIDController)
	adminRoutes.POST("/vaccine/:id", controller.UpdateVaccineController)
	adminRoutes.DELETE("/vaccine/:id", controller.DeleteVaccineIDController)
}

func RegisterStokVaccineGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoStokVaccine.NewVaccineHospitalRepository(db)

	svc := serviceStokVaccine.NewServiceVaccineHospital(repo, conf)

	controller := handlerVaccineStok.EchoControllerVaccineHospital{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/stok/create", controller.CreateStokHandler)
	adminRoutes.GET("/stok/:hospital_id", controller.GetStokByHospitalController)
	adminRoutes.GET("/stok/:hospital_id/:vaccine_id", controller.GetStokByHospitalVaccineIDController)
	adminRoutes.POST("/stok/update", controller.UpdateVaccineStokController)
	adminRoutes.DELETE("/stok/delete", controller.DeleteHospitalVaccineIDController)
}
