package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterRepositoryVaccineTransaction
}

func (s *svcVaccine) CreateVaccinesTransactionService(transaction model.VaccineTransactions) error {
	return s.repo.CreateVaccinesTransaction(transaction)
}

func (s *svcVaccine) UpdateTransactionByIDService(hospital_id, vaccine_id int, transaction model.VaccineTransactions) error {
	return s.repo.UpdateTransactionByID(hospital_id, vaccine_id, transaction)
}

func (s *svcVaccine) GetAllTransactionByHospitalService(hospital_id int) (transaction []model.VaccineTransactions, err error) {

	return s.repo.GetAllTransactionByHospital(hospital_id)
}

func (s *svcVaccine) GetTransactionByHospitalVaccineService(hospital_id, vaccine_id int) (transaction model.VaccineTransactions, err error) {
	return s.repo.GetTransactionByHospitalVaccine(hospital_id, vaccine_id)
}

func (s *svcVaccine) DeleteVaccineTransactionByIDService(hospital_id, vaccine_id int) error {
	return s.repo.DeleteVaccineTransactionByID(hospital_id, vaccine_id)
}

func NewServiceVaccineTransactions(repo domain.AdapterRepositoryVaccineTransaction, c config.Config) domain.AdapterServiceVaccineTransaction {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}
