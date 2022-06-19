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

func (r *repository) CreateHospitals(hospital model.Hospitals) error {
	res := r.DB.Create(&hospital)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}


func (r *repository) GetAllHospitals() []model.Hospitals {
	hospitals := []model.Hospitals{}
	r.DB.Find(&hospitals)

	return hospitals
}

func (r *repository) GetHospitalByID(id int) (hospital model.Hospitals, err error) {
	res := r.DB.Where("id = ?", id).Find(&hospital)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateHospitalByID(id int, hospital model.Hospitals) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&hospital)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteHospitalByID(id int) error {
	res := r.DB.Delete(&model.Hospitals{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

// By name

func (r *repository) GetHospitalByName(name string) (hospital model.Hospitals, err error) {
	res := r.DB.Where("name = ?", name).Find(&hospital)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

// func (r *repository) UpdateHospitalByName(name string, hospital model.Hospitals) error {
// 	res := r.DB.Where("name = ?", name).UpdateColumns(&hospital)
// 	if res.RowsAffected < 1 {
// 		return fmt.Errorf("error update")
// 	}

// 	return nil
// }

// func (r *repository) DeleteHospitalByName(name string) error {
// 	res := r.DB.Delete(&model.Hospitals{
// 		Name: name,
// 	})

// 	if res.RowsAffected < 1 {
// 		return fmt.Errorf("error delete")
// 	}

// 	return nil
// }

func NewHospitalRepository(db *gorm.DB) domain.AdapterRepositoryHospital {
	return &repository{
		DB: db,
	}
}
