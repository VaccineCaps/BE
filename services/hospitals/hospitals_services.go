package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcHospitals struct {
	c    config.Config
	repo domain.AdapterRepositoryHospital
}

func (s *svcHospitals) CreateHospitalService(Hospitals model.Hospitals) error {
	return s.repo.CreateHospitals(Hospitals)
}

func (s *svcHospitals) UpdateHospitalService(id int, Hospitals model.Hospitals) error {
	return s.repo.UpdateHospitalByID(id, Hospitals)
}

func (s *svcHospitals) GetAllHospitalsService() []model.Hospitals {
	return s.repo.GetAllHospitals()
}

func (s *svcHospitals) GetHospitalsByID(id int) (model.Hospitals, error) {
	return s.repo.GetHospitalByID(id)
}

// func (s *svcHospitals) GetHospitalsByName(name string) (model.Hospitals, error) {
// 	return s.repo.GetHospitalByName(name)
// }

func (s *svcHospitals) DeleteHospitalByID(id int) error {
	return s.repo.DeleteHospitalByID(id)
}

func NewServiceHospitals(repo domain.AdapterRepositoryHospital, c config.Config) domain.AdapterServiceHospital {
	return &svcHospitals{
		repo: repo,
		c:    c,
	}
}
