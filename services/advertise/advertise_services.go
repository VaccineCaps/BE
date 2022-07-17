package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcAdvertise struct {
	c    config.Config
	repo domain.AdapterRepositoryAdvertise
}

func (s *svcAdvertise) CreateAdvertiseService(Advertise model.Advertise) error {
	return s.repo.CreateAdvertise(Advertise)
}

func (s *svcAdvertise) UpdateAdvertiseService(id int, Advertise model.Advertise) error {
	return s.repo.UpdateAdvertiseByID(id, Advertise)
}

func (s *svcAdvertise) GetAllAdvertiseService() []model.Advertise {
	return s.repo.GetAllAdvertise()
}

func (s *svcAdvertise) GetAdvertiseByID(id int) (model.Advertise, error) {
	return s.repo.GetAdvertiseByID(id)
}

func (s *svcAdvertise) DeleteAdvertiseByID(id int) error {
	return s.repo.DeleteAdvertiseByID(id)
}

func NewServiceAdvertise(repo domain.AdapterRepositoryAdvertise, c config.Config) domain.AdapterServiceAdvertise {
	return &svcAdvertise{
		repo: repo,
		c:    c,
	}
}
