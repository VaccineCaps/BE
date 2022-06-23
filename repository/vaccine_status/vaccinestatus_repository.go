package repository

import (
	"fmt"

	"BE/domain"
	"BE/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateVStatus(vStatus model.VaccineStatus) error {
	res := r.DB.Create(&vStatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllVStatus() []model.VaccineStatus {
	allVStatus := []model.VaccineStatus{}
	r.DB.Find(&allVStatus)

	return allVStatus
}

func (r *repository) GetVStatusByID(id int) (vStatus model.VaccineStatus, err error) {
	res := r.DB.Where("id = ?", id).Find(&vStatus)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateVStatusByID(id int, vStatus model.VaccineStatus) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vStatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}


func (r *repository) DeleteVStatusByID(id int) error {
	res := r.DB.Delete(&model.VaccineStatus{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVStatusRepository(db *gorm.DB) domain.AdapterRepositoryVStatus {
	return &repository{
		DB: db,
	}
}
