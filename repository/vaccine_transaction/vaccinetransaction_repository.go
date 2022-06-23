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

func (r *repository) CreateVaccinesTransaction(transaction model.VaccineTransactions) error {

	res := r.DB.Create(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) UpdateTransactionByID(hospital_id, vaccine_id int, transaction model.VaccineTransactions) error {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).UpdateColumns(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) GetAllTransactionByHospital(hospital_id int) (transaction []model.VaccineTransactions, err error) {
	res := r.DB.Where("hospital_id = ? ", hospital_id).Find(&transaction)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetTransactionByHospitalVaccine(hospital_id, vaccine_id int) (transaction model.VaccineTransactions, err error) {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).Find(&transaction)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteVaccineTransactionByID(hospital_id, vaccine_id int) error {
	res := r.DB.Delete(&model.VaccineTransactions{
		HospitalId: hospital_id,
		VaccineId:  vaccine_id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVaccineTransactionRepository(db *gorm.DB) domain.AdapterRepositoryVaccineTransaction {
	return &repository{
		DB: db,
	}
}
