package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcCity struct {
	c    config.Config
	repo domain.AdapterRepositoryCity
}

func (s *svcCity) CreateCityService(City model.Cities) error {
	return s.repo.CreateCities(City)
}

func (s *svcCity) GetAllCityService() []model.Cities {
	return s.repo.GetAllCity()
}

func (s *svcCity) GetCityByID(id int) (model.Cities, error) {
	return s.repo.GetCityByID(id)
}


func (s *svcCity) DeleteCityByID(id int) error {
	return s.repo.DeleteCityByID(id)
}


func NewServiceCity(repo domain.AdapterRepositoryCity, c config.Config) domain.AdapterServiceCity {
	return &svcCity{
		repo: repo,
		c:    c,
	}
}
