package routes

import (
	config "BE/configs"
	"BE/database"
	m "BE/helper/middleware"

	"github.com/labstack/echo/v4/middleware"

	// Deklarasi Handler
	handlerAdvertise "BE/handler/advertise"
	handlerBooking "BE/handler/booking"
	handlerCertificate "BE/handler/certificate"
	handlerCities "BE/handler/cities"
	handlerDetailBook "BE/handler/detailbook"
	handlerHospital "BE/handler/hospital"
	handlerNews "BE/handler/news"
	handlerOP "BE/handler/otherperson"
	handlerProvince "BE/handler/province"
	handlerRole "BE/handler/role"
	handlerSession "BE/handler/session"
	handlerUser "BE/handler/user"
	handlerVaccine "BE/handler/vaccine"
	handlerTransactionIn "BE/handler/vaccine_transaction_in"
	handlerTransactionOut "BE/handler/vaccine_transaction_out"
	handlerVaccineStok "BE/handler/vaccinehospital"

	// Deklarasi repository
	repoAdvertise "BE/repository/advertise"
	repoBooking "BE/repository/booking"
	repoCertificate "BE/repository/certificate"
	repoCity "BE/repository/cities"
	repoDetailBook "BE/repository/detailbook"
	repoHospital "BE/repository/hospital"
	repoNews "BE/repository/news"
	repoOP "BE/repository/otherperson"
	repoProvince "BE/repository/province"
	repoRole "BE/repository/role"
	repoSession "BE/repository/session"
	repoUser "BE/repository/user"
	repoVaccine "BE/repository/vaccine"
	repoTransactionIn "BE/repository/vaccine_transaction_in"
	repoTransactionOut "BE/repository/vaccine_transaction_out"
	repoStokVaccine "BE/repository/vaccinehospital"

	// Deklarasi Services
	serviceAdvertise "BE/services/advertise"
	serviceBooking "BE/services/booking"
	serviceCertificate "BE/services/certificate"
	serviceCity "BE/services/cities"
	serviceDetailBook "BE/services/detailbook"
	serviceHospital "BE/services/hospitals"
	serviceNews "BE/services/news"
	serviceOP "BE/services/otherperson"
	serviceProvince "BE/services/province"
	serviceRole "BE/services/role"
	serviceSession "BE/services/session"
	serviceUser "BE/services/user"
	serviceVaccine "BE/services/vaccine"
	serviceTransactionIn "BE/services/vaccine_transaction_in"
	serviceTransactionOut "BE/services/vaccine_transaction_out"
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

	e.POST("admin/register", controller.RegisterHandler)
	e.POST("admin/login", controller.LoginHandler)

	e.POST("user/register", controller.RegisterHandler)
	e.POST("user/login", controller.LoginHandler)

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	// adminRoutes.POST("/register", controller.RegisterHandler)
	// adminRoutes.POST("/login", controller.LoginHandler)
	adminRoutes.GET("/users", controller.GetUsersController)
	adminRoutes.GET("/users/:id", controller.GetUserController)
	adminRoutes.POST("/users/:id", controller.UpdateUserController)
	adminRoutes.DELETE("/users/:id", controller.DeleteUserController)

	// userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	// userRoutes.POST("/register", controller.RegisterHandler)
	// userRoutes.POST("/login", controller.LoginHandler)
}

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoRole.NewRoleRepository(db)

	svc := serviceRole.NewServiceRole(repo, conf)

	controller := handlerRole.EchoControllerRole{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/hospitals", controller.CreateHospitalController)
	adminRoutes.GET("/hospitals", controller.GetHospitalController)
	adminRoutes.GET("/hospitals/:id", controller.GetHospitalIDController)
	adminRoutes.PUT("/hospitals/:id", controller.UpdateHospitalController)
	adminRoutes.DELETE("/hospitals/:id", controller.DeleteHospitalController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/news", controller.CreateNewsController)
	adminRoutes.GET("/news", controller.GetNewsController)
	adminRoutes.GET("/news/:id", controller.GetNewsIDController)
	adminRoutes.PUT("/news/:id", controller.UpdateNewsController)
	adminRoutes.DELETE("/news/:id", controller.DeleteNewsController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	userRoutes.GET("/news", controller.GetNewsController)
	userRoutes.GET("/news/:id", controller.GetNewsIDController)

}

func RegisterOPsGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoOP.NewOtherRepository(db)

	svc := serviceOP.NewServiceOtherPerson(repo, conf)

	controller := handlerOP.EchoControllerOther{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/others", controller.CreateOtherController)
	adminRoutes.GET("/others", controller.GetOtherController)
	adminRoutes.GET("/others/:id", controller.GetOtherIDController)
	adminRoutes.PUT("/others/:id", controller.UpdateOtherController)
	adminRoutes.DELETE("/others/:id", controller.DeleteOtherController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())

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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
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

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/session", controller.CreateSessionHandler)
	adminRoutes.GET("/session/:hospital_id", controller.GetSessionByHospitalController)
	adminRoutes.GET("/session/:hospital_id/:vaccine_id", controller.GetSessionByHospitalVaccineIDController)
	adminRoutes.POST("/session/:hospital_id/:vaccine_id", controller.UpdateVaccineSessionController)
	adminRoutes.DELETE("/session/:hospital_id/:vaccine_id", controller.DeleteSessionIDController)
}

func RegisterVaccineTransactionInGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoTransactionIn.NewVaccineTransactionRepository(db)

	svc := serviceTransactionIn.NewServiceVaccineTransactionsIn(repo, conf)

	controller := handlerTransactionIn.EchoControllerVaccineTransaction{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/transactionin", controller.CreateTransactionHandler)
	adminRoutes.GET("/transactionin/:id", controller.GetTransactionByIDController)
	adminRoutes.GET("/transactionin", controller.GetAllTransactionInController)
	adminRoutes.POST("/transactionin/:id", controller.UpdateVaccineTransactionController)
	adminRoutes.DELETE("/transactionin/:id", controller.DeleteVaccineTransactionIDController)
}

func RegisterVaccineTransactionOutGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoTransactionOut.NewVaccineTransactionRepository(db)

	svc := serviceTransactionOut.NewServiceVaccineTransactionsOut(repo, conf)

	controller := handlerTransactionOut.EchoControllerVaccineTransaction{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/transactionout", controller.CreateTransactionHandler)
	adminRoutes.GET("/transactionout/:id", controller.GetTransactionByIDController)
	adminRoutes.GET("/transactionout", controller.GetAllTransactionInController)
	adminRoutes.POST("/transactionout/:id", controller.UpdateVaccineTransactionController)
	adminRoutes.DELETE("/transactionout/:id", controller.DeleteVaccineTransactionIDController)
}

func RegisterCertificateGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoCertificate.NewCertificateRepository(db)

	svc := serviceCertificate.NewServiceCertificate(repo, conf)

	controller := handlerCertificate.EchoControllerCertificate{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/certificate", controller.CreateCertificateController)
	adminRoutes.GET("/certificate", controller.GetAllCertificateController)
	adminRoutes.GET("/certificate/:id", controller.GetCertificateIDController)
	adminRoutes.PUT("/certificate/:id", controller.UpdateCertificateController)
	adminRoutes.DELETE("/certificate/:id", controller.DeleteCertificateIDController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	userRoutes.GET("/certificate", controller.GetAllCertificateController)
	userRoutes.GET("/certificate/:id", controller.GetCertificateIDController)
}

func RegisterBookingGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoBooking.NewBookingRepository(db)
	repoSession := repoSession.NewSessionRepository(db)

	svc := serviceBooking.NewServiceBooking(repo, repoSession, conf)

	controller := handlerBooking.EchoControllerBooking{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/booking", controller.CreateBookingHandler)
	adminRoutes.GET("/booking", controller.GetAllBookingController)
	// adminRoutes.GET("/booking/:user_id", controller.GetBookingByUserController)
	adminRoutes.GET("/booking/:session_id", controller.GetBookingBySessionController)
	adminRoutes.DELETE("/booking/:user_id/:hospital_id/:session_id/:vaccinestatus_id", controller.DeleteBookingController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	userRoutes.GET("/booking/:user_id", controller.GetBookingByUserController)

}

func RegisterBookingDetailGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoDetailBook.NewBookingDetailRepository(db)

	svc := serviceDetailBook.NewServiceBookingDetail(repo, conf)

	controller := handlerDetailBook.EchoControllerBookingDetail{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/detail", controller.CreateBookingDetailHandler)
	adminRoutes.GET("/detail", controller.GetAllBookingDetailController)
	// adminRoutes.GET("/detail/:id", controller.GetBookingDetailByIDController)
	adminRoutes.GET("/detail/:user_id", controller.GetBookingDetailByUserController)
	// adminRoutes.GET("/detail/:otherperson_id", controller.GetBookingDetailByOPController)
	// adminRoutes.GET("/detail/:booking_id", controller.GetBookingDetailByBookingController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	userRoutes.GET("/detail/:user_id", controller.GetBookingDetailByUserController)
}

func RegisterAdvertiseGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repoAdvertise.NewAdvertiseRepository(db)

	svc := serviceAdvertise.NewServiceAdvertise(repo, conf)

	controller := handlerAdvertise.EchoControllerAdvertise{
		Svc: svc,
	}

	adminRoutes := e.Group("admin", m.CheckTokenAdmin, middleware.CORS())
	adminRoutes.POST("/advertise", controller.CreateAdvertiseController)
	adminRoutes.GET("/advertise", controller.GetAdvertiseController)
	adminRoutes.GET("/advertise/:id", controller.GetAdvertiseIDController)
	adminRoutes.PUT("/advertise/:id", controller.UpdateAdvertiseController)
	adminRoutes.DELETE("/advertise/:id", controller.DeleteAdvertiseController)

	userRoutes := e.Group("user", m.CheckTokenUser, middleware.CORS())
	userRoutes.GET("/advertise", controller.GetAdvertiseController)
	userRoutes.GET("/advertise/:id", controller.GetAdvertiseIDController)
}
