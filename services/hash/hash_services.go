package service

import (
	config "BE/configs"
	"BE/domain"
)

type svcHash struct {
	c    config.Config
	repo domain.AdapterRepositoryHash
}

func (s *svcHash) HashPasswordService(password string) (string, error) {
	return s.repo.HashPassword(password)
}

func (s *svcHash) CheckPasswordHashService(password, hashed string) bool {
	return s.repo.CheckPasswordHash(password, hashed)
}

func NewServiceHAsh(repo domain.AdapterRepositoryHash, c config.Config) domain.AdapterServiceHash {
	return &svcHash{
		repo: repo,
		c:    c,
	}
}
