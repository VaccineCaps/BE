package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcSession struct {
	c    config.Config
	repo domain.AdapterRepositorySession
}

func (s *svcSession) CreateSessionService(session model.Session) error {
	return s.repo.CreateSession(session)
}

func (s *svcSession) UpdateSessionService(hospital_id, vaccine_id int, session model.Session) error {
	return s.repo.UpdateSessionByID(hospital_id, vaccine_id, session)
}

func (s *svcSession) GetAllSessionByHospitalService(hospital_id int) (session []model.Session, err error) {

	return s.repo.GetAllSession(hospital_id)
}

func (s *svcSession) GetSessionByHospitalVaccineService(hospital_id, vaccine_id int) (session model.Session, err error) {
	return s.repo.GetSessionByHospitalVaccine(hospital_id, vaccine_id)
}

func (s *svcSession) DeleteSessionByIDService(hospital_id, vaccine_id int) error {
	return s.repo.DeleteSessionByID(hospital_id, vaccine_id)
}

func NewServiceSession(repo domain.AdapterRepositorySession, c config.Config) domain.AdapterServiceSession {
	return &svcSession{
		repo: repo,
		c:    c,
	}
}
