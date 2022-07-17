package domain

import (
	"BE/model"
)

// ================================================ Repository
type AdapterRepositoryHash interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hashed string) bool
}

type AdapterRepositoryUser interface {

	// untuk user
	Login(email, password string) (credential model.User, err error)
	CreateUsers(user model.User) error
	GetAll() []model.User
	GetOneByID(id int) (user model.User, err error)
	GetOneByEmail(email string) (user model.User, err error)
	UpdateOneByID(id int, user model.User) error
	DeleteByID(id int) error
}

type AdapterRepositoryRole interface {

	// untuk Role
	CreateRoles(role model.Role) error
	GetAllRole() []model.Role
	GetRoleByID(id int) (role model.Role, err error)
	// GetRoleByName(name string) (role model.Role, err error)
	DeleteRoleByID(id int) error
	// DeleteRoleByName(name string) error
}

type AdapterRepositoryProvince interface {

	// untuk Province
	CreateProvinces(province model.Provinces) error
	GetAllProvince() []model.Provinces
	GetProvinceByID(id int) (province model.Provinces, err error)
	DeleteProvinceByID(id int) error
}

type AdapterRepositoryCity interface {

	// untuk City
	CreateCities(City model.Cities) error
	GetAllCity() []model.Cities
	GetCityByID(id int) (City model.Cities, err error)
	DeleteCityByID(id int) error
}

type AdapterRepositoryHospital interface {

	// untuk hospital
	CreateHospitals(hospital model.Hospitals) error
	GetAllHospitals() []model.Hospitals
	// GetHospitalByName(name string) (hospital model.Hospitals, err error)
	GetHospitalByID(id int) (hospital model.Hospitals, err error)
	UpdateHospitalByID(id int, hospital model.Hospitals) error
	DeleteHospitalByID(id int) error
	// UpdateHospitalByName(name string, hospital model.Hospitals) error
	// DeleteHospitalByName(name string) error
}

type AdapterRepositoryNews interface {
	CreateNews(news model.News) error
	GetAllNews() []model.News
	GetNewsByID(id int) (news model.News, err error)
	UpdateNewsByID(id int, news model.News) error
	DeleteNewsByID(id int) error
}

type AdapterRepositoryOther interface {

	//otherperson
	CreateOtherPerson(Other model.OtherPerson) error
	GetAllOtherPerson() []model.OtherPerson
	GetOtherByID(id int) (Other model.OtherPerson, err error)
	UpdateOtherByID(id int, Other model.OtherPerson) error
	DeleteOtherByID(id int) error
}

type AdapterRepositoryVaccine interface {
	// Vaccine
	CreateVaccines(vaccine model.Vaccines) error
	GetAllVaccine() []model.Vaccines
	GetVaccineByID(id int) (vaccine model.Vaccines, err error)
	UpdateVaccineByID(id int, vaccine model.Vaccines) error
	DeleteVaccineByID(id int) error
}

type AdapterRepositoryVaccineHospital interface {
	// Hospital Stok Vaccine
	CreateStokVaccineHospital(stok model.VaccineHospitals) error
	UpdateStokByID(hospital_id, vaccine_id int, stok model.VaccineHospitals) error
	GetAllStokByHospital(hospital_id int) (stok []model.VaccineHospitals, err error)
	GetStokByHospitalVaccine(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error)
	DeleteVaccineStokByID(hospital_id, vaccine_id int) error
}

type AdapterRepositorySession interface {
	CreateSession(session model.Session) error
	UpdateSessionByID(hospital_id, vaccine_id int, session model.Session) error
	GetAllSession(hospital_id int) (session []model.Session, err error)
	GetSessionByHospitalVaccine(hospital_id, vaccine_id int) (session model.Session, err error)
	DeleteSessionByID(hospital_id, vaccine_id int) error
}

type AdapterRepositoryCertificate interface {
	CreateCertificate(vStatus model.Certificate) error
	GetAllCertificate() []model.Certificate
	GetCertificateByID(id int) (vStatus model.Certificate, err error)
	UpdateCertificateByID(id int, vStatus model.Certificate) error
	DeleteCertificateByID(id int) error
}

type AdapterRepositoryBooking interface {
	CreateBooking(booking model.Booking) error
	GetAllBooking() []model.Booking
	GetAllBookingByUsers(user_id int) (booking []model.Booking, err error)
	GetAllBookingBySessions(session_id int) (booking []model.Booking, err error)
	GetBookingsByID(id int) (booking model.Booking, err error)
	DeleteBookingByID(user_id, hospital_id, session_id, vaccinestatus_id int) error
}

type AdapterRepositoryBookingDetail interface {
	CreateBookingDetail(book model.BookingDetail) error
	GetDetailByID(id int) (booking model.BookingDetail, err error)
	GetAllBookingDetail() []model.BookingDetail
	GetDetailByUsers(user_id int) (booking []model.BookingDetail, err error)
	GetDetailByOtherPerson(otherperson_id int) (booking []model.BookingDetail, err error)
	GetDetailByBookings(booking_id int) (booking []model.BookingDetail, err error)
}

type AdapterRepositoryVaccineTransactionIn interface {
	CreateVaccinesTransaction(transaction model.VaccineTransactionsIn) error
	UpdateTransactionByID(id int, transaction model.VaccineTransactionsIn) error
	GetTransactionByID(id int) (transaction []model.VaccineTransactionsIn, err error)
	GetAllTransaction() []model.VaccineTransactionsIn
	DeleteVaccineTransactionByID(id int) error
}

type AdapterRepositoryVaccineTransactionOut interface {
	CreateVaccinesTransaction(transaction model.VaccineTransactionsOut) error
	UpdateTransactionByID(id int, transaction model.VaccineTransactionsOut) error
	GetTransactionByID(id int) (transaction []model.VaccineTransactionsOut, err error)
	GetAllTransaction() []model.VaccineTransactionsOut
	DeleteVaccineTransactionByID(id int) error
}

type AdapterRepositoryAdvertise interface {
	CreateAdvertise(advertise model.Advertise) error
	GetAllAdvertise() []model.Advertise
	GetAdvertiseByID(id int) (advertise model.Advertise, err error)
	UpdateAdvertiseByID(id int, advertise model.Advertise) error
	DeleteAdvertiseByID(id int) error
}

// ================================================ Service
type AdapterServiceHash interface {
	HashPasswordService(password string) (string, error)
	CheckPasswordHashService(password, hashed string) bool
}

type AdapterServiceUser interface {

	//untuk user
	CreateUserService(user model.User) error
	UpdateUserService(id, idToken int, user model.User) error
	GetAllUsersService() []model.User
	GetUserByID(id int) (model.User, error)
	LoginUser(email, password string) (string, int)
	DeleteByID(id int) error
}

type AdapterServiceRole interface {

	//untuk role
	CreateRoleService(role model.Role) error
	GetAllRoleService() []model.Role
	GetRoleByID(id int) (model.Role, error)
	// GetRoleByName(name string) (model.Role, error)
	DeleteRoleByID(id int) error
	// DeleteRoleByName(name string) error
}

type AdapterServiceProvince interface {

	//province
	CreateProvinceService(province model.Provinces) error
	GetAllProvinceService() []model.Provinces
	GetProvinceByID(id int) (model.Provinces, error)
	DeleteProvinceByID(id int) error
}

type AdapterServiceCity interface {

	//city
	CreateCityService(City model.Cities) error
	GetAllCityService() []model.Cities
	GetCityByID(id int) (model.Cities, error)
	DeleteCityByID(id int) error
}

type AdapterServiceHospital interface {

	//hospital
	CreateHospitalService(Hospitals model.Hospitals) error
	UpdateHospitalService(id int, Hospitals model.Hospitals) error
	GetAllHospitalsService() []model.Hospitals
	GetHospitalsByID(id int) (model.Hospitals, error)
	// GetHospitalsByName(name string) (model.Hospitals, error)
	DeleteHospitalByID(id int) error
}

type AdapterServiceNews interface {

	//news
	CreateNewsService(News model.News) error
	UpdateNewsService(id int, News model.News) error
	GetAllNewsService() []model.News
	GetNewsByID(id int) (model.News, error)
	DeleteNewsByID(id int) error
}

type AdapterServiceOther interface {

	//otherperson
	CreateOtherService(OtherPerson model.OtherPerson) error
	UpdateOtherService(id int, OtherPerson model.OtherPerson) error
	GetAllOtherService() []model.OtherPerson
	GetOtherByID(id int) (model.OtherPerson, error)
	DeleteOtherByID(id int) error
}

type AdapterServiceVaccine interface {
	// Vaccine
	CreateVaccineService(vaccine model.Vaccines) error
	GetAllVaccineService() []model.Vaccines
	GetVaccineByID(id int) (model.Vaccines, error)
	UpdateVaccineByID(id int, vaccine model.Vaccines) error
	DeleteVaccineByID(id int) error
}

type AdapterServiceVaccineHospital interface {
	// Hospital Stok Vaccine
	CreateStokService(stok model.VaccineHospitals) error
	UpdateVaccineStokService(hospital_id, vaccine_id int, stok model.VaccineHospitals) error
	GetAllStokByHospitalService(hospital_id int) (stok []model.VaccineHospitals, err error)
	GetStokByHospitalVaccineService(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error)
	DeleteVaccineStokByIDService(hospital_id, vaccine_id int) error
}

type AdapterServiceSession interface {
	CreateSessionService(session model.Session) error
	UpdateSessionService(hospital_id, vaccine_id int, session model.Session) error
	GetAllSessionByHospitalService(hospital_id int) (session []model.Session, err error)
	GetSessionByHospitalVaccineService(hospital_id, vaccine_id int) (session model.Session, err error)
	DeleteSessionByIDService(hospital_id, vaccine_id int) error
}

type AdapterServiceCertificate interface {
	CreateCertificateService(vStatus model.Certificate) error
	GetAllCertificateService() []model.Certificate
	GetCertificateByID(id int) (model.Certificate, error)
	UpdateCertificateService(id int, vStatus model.Certificate) error
	DeleteCertificateByID(id int) error
}

type AdapterServiceBooking interface {
	CreateBookingService(booking model.Booking) error
	GetAllBookingService() []model.Booking
	GetBookingByUserService(user_id int) (booking []model.Booking, err error)
	GetAllBookingBySessionsService(session_id int) (booking []model.Booking, err error)
	GetBookingsByIDService(id int) (booking model.Booking, err error)
	DeleteBookingByIDService(user_id, hospital_id, session_id, vaccinestatus_id int) error
}

type AdapterServiceBookingDetail interface {
	CreateBookingDetailService(booking model.BookingDetail) error
	GetAllDetailService() []model.BookingDetail
	GetBookingsByIDService(id int) (booking model.BookingDetail, err error)
	GetBookingDetailByUserService(user_id int) (booking []model.BookingDetail, err error)
	GetBookingDetailByOtherPersonService(otherperson_id int) (booking []model.BookingDetail, err error)
	GetBookingDetailByBookingsService(booking_id int) (booking []model.BookingDetail, err error)
}

type AdapterServiceVaccineTransactionIn interface {
	CreateVaccinesTransactionService(transaction model.VaccineTransactionsIn) error
	UpdateTransactionByIDService(id int, transaction model.VaccineTransactionsIn) error
	GetTransactionByIDService(id int) (transaction []model.VaccineTransactionsIn, err error)
	GetAllTransactionService() []model.VaccineTransactionsIn
	DeleteVaccineTransactionByIDService(id int) error
}

type AdapterServiceVaccineTransactionOut interface {
	CreateVaccinesTransactionService(transaction model.VaccineTransactionsOut) error
	UpdateTransactionByIDService(id int, transaction model.VaccineTransactionsOut) error
	GetTransactionByIDService(id int) (transaction []model.VaccineTransactionsOut, err error)
	GetAllTransactionService() []model.VaccineTransactionsOut
	DeleteVaccineTransactionByIDService(id int) error
}

type AdapterServiceAdvertise interface {
	CreateAdvertiseService(Advertise model.Advertise) error
	UpdateAdvertiseService(id int, Advertise model.Advertise) error
	GetAllAdvertiseService() []model.Advertise
	GetAdvertiseByID(id int) (model.Advertise, error)
	DeleteAdvertiseByID(id int) error
}
