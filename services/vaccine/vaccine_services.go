package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterRepositoryVaccine
}

func (s *svcVaccine) CreateVaccineService(vaccine model.Vaccines) error {
	return s.repo.CreateVaccines(vaccine)
}

func (s *svcVaccine) GetAllVaccineService() []model.Vaccines {
	return s.repo.GetAllVaccine()
}

func (s *svcVaccine) GetVaccineByID(id int) (model.Vaccines, error) {
	return s.repo.GetVaccineByID(id)
}

func (s *svcVaccine) UpdateVaccineByID(id int, vaccine model.Vaccines) error {
	return s.repo.UpdateVaccineByID(id, vaccine)
}

func (s *svcVaccine) DeleteVaccineByID(id int) error {
	return s.repo.DeleteVaccineByID(id)
}

func NewServiceVaccine(repo domain.AdapterRepositoryVaccine, c config.Config) domain.AdapterServiceVaccine {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}
