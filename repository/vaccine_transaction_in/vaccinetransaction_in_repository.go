package repository

import (
	"BE/domain"
	"BE/model"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateVaccinesTransaction(transaction model.VaccineTransactionsIn) error {

	res := r.DB.Create(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) UpdateTransactionByID(id int, transaction model.VaccineTransactionsIn) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) GetTransactionByID(id int) (transaction []model.VaccineTransactionsIn, err error) {
	res := r.DB.Where("id = ? ", id).Find(&transaction)
	fmt.Println(res.Error)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetAllTransaction() []model.VaccineTransactionsIn {
	transaction_in := []model.VaccineTransactionsIn{}
	r.DB.Find(&transaction_in)

	return transaction_in
}

func (r *repository) DeleteVaccineTransactionByID(id int) error {
	res := r.DB.Delete(&model.VaccineTransactionsIn{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVaccineTransactionRepository(db *gorm.DB) domain.AdapterRepositoryVaccineTransactionIn {
	return &repository{
		DB: db,
	}
}
