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

func (r *repository) CreateProvinces(province model.Provinces) error {
	res := r.DB.Create(&province)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllProvince() []model.Provinces {
	allprovince := []model.Provinces{}
	r.DB.Find(&allprovince)

	return allprovince
}

func (r *repository) GetProvinceByID(id int) (province model.Provinces, err error) {
	res := r.DB.Where("id = ?", id).Find(&province)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteProvinceByID(id int) error {
	res := r.DB.Delete(&model.Provinces{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewProvinceRepository(db *gorm.DB) domain.AdapterRepositoryProvince {
	return &repository{
		DB: db,
	}
}
