package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcCertificate struct {
	c    config.Config
	repo domain.AdapterRepositoryCertificate
}

func (s *svcCertificate) CreateCertificateService(vStatus model.Certificate) error {
	return s.repo.CreateCertificate(vStatus)
}

func (s *svcCertificate) GetAllCertificateService() []model.Certificate {
	return s.repo.GetAllCertificate()
}

func (s *svcCertificate) GetCertificateByID(id int) (model.Certificate, error) {
	return s.repo.GetCertificateByID(id)
}

func (s *svcCertificate) UpdateCertificateService(id int, vStatus model.Certificate) error{
	return s.repo.UpdateCertificateByID(id, vStatus)
}

func (s *svcCertificate) DeleteCertificateByID(id int) error {
	return s.repo.DeleteCertificateByID(id)
}

func NewServiceCertificate(repo domain.AdapterRepositoryCertificate, c config.Config) domain.AdapterServiceCertificate {
	return &svcCertificate{
		repo: repo,
		c:    c,
	}
}
