package routes

import (
	config "BE/configs"
	"BE/database"
	"BE/helper/middleware"

	// Deklarasi Handler
	handlerCities 		"BE/handler/cities"
	handlerHospital 	"BE/handler/hospital"
	handlerOP 			"BE/handler/otherperson"
	handlerVStatus		"BE/handler/vaccine_status"
	handlerNews 		"BE/handler/news"
	handlerProvince 	"BE/handler/province"
	handlerRole 		"BE/handler/role"
	handlerSession 		"BE/handler/session"
	handlerUser 		"BE/handler/user"
	handlerVaccine 		"BE/handler/vaccine"
	handlerVaccineStok 	"BE/handler/vaccinehospital"
	handlerBooking		"BE/handler/booking"
	handlerDetailBook	"BE/handler/detailbook"

	// Deklarasi repository 
	repoRole 		"BE/repository/role"
	repoUser 		"BE/repository/user"
	repoProvince 	"BE/repository/province"
	repoCity 		"BE/repository/cities"
	repoHospital 	"BE/repository/hospital"
	repoNews 		"BE/repository/news"
	repoOP			"BE/repository/otherperson"
	repoVStatus 	"BE/repository/vaccine_status"
	repoVaccine 	"BE/repository/vaccine"
	repoStokVaccine "BE/repository/vaccinehospital"
	repoSession 	"BE/repository/session"
	repoBooking 	"BE/repository/booking"
	repoDetailBook 	"BE/repository/detailbook"

	// Deklarasi Services 
	serviceRole 		"BE/services/role"
	serviceUser 		"BE/services/user"
	serviceProvince 	"BE/services/province"
	serviceCity 		"BE/services/cities"
	serviceHospital 	"BE/services/hospitals"
	serviceNews			"BE/services/news"
	serviceOP 			"BE/services/otherperson"
	serviceVStatus		"BE/services/vaccine_status"
	serviceSession 		"BE/services/session"
	serviceStokVaccine 	"BE/services/vaccinehospital"
	serviceVaccine 		"BE/services/vaccine"
	serviceBooking 		"BE/services/booking"
	serviceDetailBook	"BE/services/detailBook"

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

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.POST("/others", controller.CreateOtherController)
	userRoutes.GET("/others", controller.GetOtherController)
	userRoutes.GET("/others/:id", controller.GetOtherIDController)
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
	adminRoutes.POST("/stok", controller.CreateStokHandler)
	adminRoutes.GET("/stok/:hospital_id", controller.GetStokByHospitalController)
	adminRoutes.GET("/stok/:hospital_id/:vaccine_id", controller.GetStokByHospitalVaccineIDController)
	adminRoutes.POST("/stok/:hospital_id/:vaccine_id", controller.UpdateVaccineStokController)
	adminRoutes.DELETE("/stok/:hospital_id/:vaccine_id", controller.DeleteHospitalVaccineIDController)
}

func RegisterSessionGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoSession.NewSessionRepository(db)

	svc := serviceSession.NewServiceSession(repo, conf)

	controller := handlerSession.EchoControllerSession{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/session", controller.CreateSessionHandler)
	adminRoutes.GET("/session/:hospital_id", controller.GetSessionByHospitalController)
	adminRoutes.GET("/session/:hospital_id/:vaccine_id", controller.GetSessionByHospitalVaccineIDController)
	adminRoutes.POST("/session/:hospital_id/:vaccine_id", controller.UpdateVaccineSessionController)
	adminRoutes.DELETE("/session/:hospital_id/:vaccine_id", controller.DeleteSessionIDController)
}

func RegisterVStatusGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoVStatus.NewVStatusRepository(db)

	svc := serviceVStatus.NewServiceVStatus(repo, conf)

	controller := handlerVStatus.EchoControllerVStatus{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
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

func RegisterBookingGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoBooking.NewBookingRepository(db)

	svc := serviceBooking.NewServiceBooking(repo, conf)

	controller := handlerBooking.EchoControllerBooking{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/booking", controller.CreateBookingHandler)
	adminRoutes.GET("/booking", controller.GetAllBookingController)
	adminRoutes.GET("/booking/:user_id", controller.GetBookingByUserController)
	adminRoutes.GET("/booking/:session_id", controller.GetBookingBySessionController)
	adminRoutes.DELETE("/booking/:user_id/:hospital_id/:session_id/:vaccinestatus_id", controller.DeleteBookingController)

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.GET("/booking/:user_id", controller.GetBookingByUserController)

}

func RegisterBookingDetailGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoDetailBook.NewBookingDetailRepository(db)

	svc := serviceDetailBook.NewServiceBookingDetail(repo, conf)

	controller := handlerDetailBook.EchoControllerBookingDetail{
		Svc: svc,
	}

	adminRoutes := e.Group("admin")
	adminRoutes.Use(middleware.CheckTokenAdmin)
	adminRoutes.POST("/detail", controller.CreateBookingDetailHandler)
	adminRoutes.GET("/detail", controller.GetAllBookingDetailController)
	adminRoutes.GET("/detail/:id", controller.GetBookingDetailByIDController)
	adminRoutes.GET("/detail/:user_id", controller.GetBookingDetailByUserController)
	adminRoutes.GET("/detail/:otherperson_id", controller.GetBookingDetailByOPController)
	adminRoutes.GET("/detail/:booking_id", controller.GetBookingDetailByBookingController)

	userRoutes := e.Group("user")
	userRoutes.Use(middleware.CheckTokenUser)
	userRoutes.GET("/detail/:user_id", controller.GetBookingDetailByUserController)
}

