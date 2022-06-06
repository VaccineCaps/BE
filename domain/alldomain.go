package domain

import (
	"BE/model"
)

// ================================================ Repository

type AdapterRepositoryUser interface {

	// untuk user
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
	GetRoleByName(name string) (role model.Role, err error)
	DeleteRoleByID(id int) error
	DeleteRoleByName(id int) error
}

// ================================================ Service

type AdapterServiceUser interface{

	//untuk user
	CreateUserService(user model.User) error
	UpdateUserService(id, idToken int, user model.User) error
	GetAllUsersService() []model.User
	GetUserByID(id int) (model.User, error)
	LoginUser(email, password string) (string, int)
	DeleteByID(id int) error
}

type AdapterServiceRole interface{

	//untuk role
	CreateRoleService(role model.Role) error
	GetAllRolesService() []model.Role
	GetRoleByID(id int) (model.Role, error)
	GetRoleByName(id int) (model.Role, error)
	DeleteRoleByID(id int) error
	DeleteRoleByName(name string) error
}