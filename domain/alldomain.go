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
