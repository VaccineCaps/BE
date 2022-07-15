package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
	"time"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterRepositoryVaccineTransactionIn
}

func (s *svcVaccine) CreateVaccinesTransactionService(transaction model.VaccineTransactionsIn) error {
	transaction.Tanggal = time.Now()

	return s.repo.CreateVaccinesTransaction(transaction)
}

func (s *svcVaccine) UpdateTransactionByIDService(id int, transaction model.VaccineTransactionsIn) error {
	return s.repo.UpdateTransactionByID(id, transaction)
}

func (s *svcVaccine) GetTransactionByIDService(id int) (transaction []model.VaccineTransactionsIn, err error) {

	transaction, err = s.repo.GetTransactionByID(id)
	// fmt.Println(transaction)
	return
}

func (s *svcVaccine) GetAllTransactionService() []model.VaccineTransactionsIn {
	return s.repo.GetAllTransaction()
}

func (s *svcVaccine) DeleteVaccineTransactionByIDService(id int) error {
	return s.repo.DeleteVaccineTransactionByID(id)
}

func NewServiceVaccineTransactionsIn(repo domain.AdapterRepositoryVaccineTransactionIn, c config.Config) domain.AdapterServiceVaccineTransactionIn {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}
