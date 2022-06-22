package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcOtherPerson struct {
	c    config.Config
	repo domain.AdapterRepositoryOther
}

func (s *svcOtherPerson) CreateOtherService(OtherPerson model.OtherPerson) error {
	return s.repo.CreateOtherPerson(OtherPerson)
}

func (s *svcOtherPerson) UpdateOtherService(id int, OtherPerson model.OtherPerson) error {
	return s.repo.UpdateOtherByID(id, OtherPerson)
}

func (s *svcOtherPerson) GetAllOtherService() []model.OtherPerson {
	return s.repo.GetAllOtherPerson()
}

func (s *svcOtherPerson) GetOtherByID(id int) (model.OtherPerson, error) {
	return s.repo.GetOtherByID(id)
}

func (s *svcOtherPerson) DeleteOtherByID(id int) error {
	return s.repo.DeleteOtherByID(id)
}

func NewServiceOtherPerson(repo domain.AdapterRepositoryOther, c config.Config) domain.AdapterServiceOther {
	return &svcOtherPerson{
		repo: repo,
		c:    c,
	}
}
