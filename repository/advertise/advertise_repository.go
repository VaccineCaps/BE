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

func (r *repository) CreateAdvertise(advertise model.Advertise) error {

	res := r.DB.Create(&advertise)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllAdvertise() []model.Advertise {
	advertise := []model.Advertise{}
	r.DB.Find(&advertise)

	return advertise
}

func (r *repository) GetAdvertiseByID(id int) (advertise model.Advertise, err error) {
	res := r.DB.Where("id = ?", id).Find(&advertise)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateAdvertiseByID(id int, advertise model.Advertise) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&advertise)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteAdvertiseByID(id int) error {
	res := r.DB.Delete(&model.Advertise{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewAdvertiseRepository(db *gorm.DB) domain.AdapterRepositoryAdvertise {
	return &repository{
		DB: db,
	}
}
