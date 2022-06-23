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

type AdapterRepositoryVStatus interface{
	CreateVStatus(vStatus model.VaccineStatus) error
	GetAllVStatus() []model.VaccineStatus
	GetVStatusByID(id int) (vStatus model.VaccineStatus, err error)
	UpdateVStatusByID(id int, vStatus model.VaccineStatus) error
	DeleteVStatusByID(id int) error
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

type AdapterServiceVStatus interface{
	CreateVStatusService(vStatus model.VaccineStatus) error
	GetAllVStatusService() []model.VaccineStatus
	GetVStatusByID(id int) (model.VaccineStatus, error)
	UpdateVStatusService(id int, vStatus model.VaccineStatus) error
	DeleteVStatusByID(id int) error
}