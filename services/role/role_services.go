package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcRole struct {
	c    config.Config
	repo domain.AdapterRepositoryRole
}

func (s *svcRole) CreateRoleService(role model.Role) error {
	return s.repo.CreateRoles(role)
}

func (s *svcRole) GetAllRolesService() []model.Role {
	return s.repo.GetAllRole()
}

func (s *svcRole) GetRoleByID(id int) (model.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *svcRole) GetRoleByName(name string) (model.Role, error) {
	return s.repo.GetRoleByName(name)
}

func (s *svcRole) DeleteRoleByID(id int) error {
	return s.repo.DeleteRoleByID(id)
}

func (s *svcRole) DeleteRoleByName(name string) error {
	return s.repo.DeleteRoleByName(name)
}

func NewServiceRole(repo domain.AdapterRepositoryRole, c config.Config) domain.AdapterServiceRole {
	return &svcRole{
		repo: repo,
		c:    c,
	}
}
