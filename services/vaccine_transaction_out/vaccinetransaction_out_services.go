package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
	"time"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterRepositoryVaccineTransactionOut
}

func (s *svcVaccine) CreateVaccinesTransactionService(transaction model.VaccineTransactionsOut) error {
	transaction.Tanggal = time.Now()

	return s.repo.CreateVaccinesTransaction(transaction)
}

func (s *svcVaccine) UpdateTransactionByIDService(id int, transaction model.VaccineTransactionsOut) error {
	return s.repo.UpdateTransactionByID(id, transaction)
}

func (s *svcVaccine) GetTransactionByIDService(id int) (transaction []model.VaccineTransactionsOut, err error) {

	transaction, err = s.repo.GetTransactionByID(id)
	// fmt.Println(transaction)
	return
}

func (s *svcVaccine) GetAllTransactionService() []model.VaccineTransactionsOut {
	return s.repo.GetAllTransaction()
}

func (s *svcVaccine) DeleteVaccineTransactionByIDService(id int) error {
	return s.repo.DeleteVaccineTransactionByID(id)
}

func NewServiceVaccineTransactionsOut(repo domain.AdapterRepositoryVaccineTransactionOut, c config.Config) domain.AdapterServiceVaccineTransactionOut {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}
