package service

import (
	"fmt"
	"net/http"

	"BE/config"
	"BE/domain"
	"BE/helper"
	"BE/model"
)

type svcRole struct {
	c config.Config
	repo domain.AdapterRepositoryRole
}

func (s *svcRole) CreateRoleService(role model.Role) error {
	return s.repo.Createroles(role)
}

func (s *svcRole) GetAllRolesService() []model.Role {
	return s.repo.GetAllRole()
}

func (s *svcRole) GetRoleByID(id int) (model.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *svcRole) GetRoleByName(id int) (model.Role, error) {
	return s.repo.GetRoleByName(id)
}

func (s *svcRole) DeleteRoleByID(id int) error {
	return s.repo.DeleteRoleByID(id)
}

func (s *svcRole) DeleteRoleByName(name string) error {
	return s.repo.DeleteRoleByName(name string)
}

func NewServiceRole(repo domain.AdapterRepositoryRole, c config.Config) domain.AdapterServiceRole {
	return &svcRole{
		repo: repo,
		c: c,
	}
}