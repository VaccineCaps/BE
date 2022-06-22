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

func (r *repository) CreateCities(City model.Cities) error {
	res := r.DB.Create(&City)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllCity() []model.Cities {
	allCity := []model.Cities{}
	r.DB.Find(&allCity)

	return allCity
}

func (r *repository) GetCityByID(id int) (City model.Cities, err error) {
	res := r.DB.Where("id = ?", id).Find(&City)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteCityByID(id int) error {
	res := r.DB.Delete(&model.Cities{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewCityRepository(db *gorm.DB) domain.AdapterRepositoryCity {
	return &repository{
		DB: db,
	}
}
