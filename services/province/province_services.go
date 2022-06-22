package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcProvince struct {
	c    config.Config
	repo domain.AdapterRepositoryProvince
}

func (s *svcProvince) CreateProvinceService(province model.Provinces) error {
	return s.repo.CreateProvinces(province)
}

func (s *svcProvince) GetAllProvinceService() []model.Provinces {
	return s.repo.GetAllProvince()
}

func (s *svcProvince) GetProvinceByID(id int) (model.Provinces, error) {
	return s.repo.GetProvinceByID(id)
}


func (s *svcProvince) DeleteProvinceByID(id int) error {
	return s.repo.DeleteProvinceByID(id)
}


func NewServiceProvince(repo domain.AdapterRepositoryProvince, c config.Config) domain.AdapterServiceProvince {
	return &svcProvince{
		repo: repo,
		c:    c,
	}
}
