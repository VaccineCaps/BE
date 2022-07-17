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

func (r *repository) CreateVaccines(vaccine model.Vaccines) error {
	res := r.DB.Create(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllVaccine() []model.Vaccines {
	allvaccine := []model.Vaccines{}
	r.DB.Find(&allvaccine)

	return allvaccine
}

func (r *repository) GetVaccineByID(id int) (vaccine model.Vaccines, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccine)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateVaccineByID(id int, vaccine model.Vaccines) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteVaccineByID(id int) error {
	res := r.DB.Delete(&model.Vaccines{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVaccineRepository(db *gorm.DB) domain.AdapterRepositoryVaccine {
	return &repository{
		DB: db,
	}
}
