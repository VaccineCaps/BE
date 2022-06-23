package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcVStatus struct {
	c    config.Config
	repo domain.AdapterRepositoryVStatus
}

func (s *svcVStatus) CreateVStatusService(vStatus model.VaccineStatus) error {
	return s.repo.CreateVStatus(vStatus)
}

func (s *svcVStatus) GetAllVStatusService() []model.VaccineStatus {
	return s.repo.GetAllVStatus()
}

func (s *svcVStatus) GetVStatusByID(id int) (model.VaccineStatus, error) {
	return s.repo.GetVStatusByID(id)
}

func (s *svcVStatus) UpdateVStatusService(id int, vStatus model.VaccineStatus) error{
	return s.repo.UpdateVStatusByID(id, vStatus)
}

func (s *svcVStatus) DeleteVStatusByID(id int) error {
	return s.repo.DeleteVStatusByID(id)
}

func NewServiceVStatus(repo domain.AdapterRepositoryVStatus, c config.Config) domain.AdapterServiceVStatus {
	return &svcVStatus{
		repo: repo,
		c:    c,
	}
}
