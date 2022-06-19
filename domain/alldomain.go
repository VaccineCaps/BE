package domain

import (
	"BE/model"
)

// ================================================ Repository

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

// ================================================ Service

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

type AdapterServiceHospital interface{

	//hospital
	CreateHospitalService(Hospitals model.Hospitals) error 
	UpdateHospitalService(id int, Hospitals model.Hospitals) error 
	GetAllHospitalsService() []model.Hospitals 
	GetHospitalsByID(id int) (model.Hospitals, error) 
	DeleteHospitalByID(id int) error 

}