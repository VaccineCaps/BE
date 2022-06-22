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

func (r *repository) CreateOtherPerson(Other model.OtherPerson) error {
	res := r.DB.Create(&Other)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}


func (r *repository) GetAllOtherPerson() []model.OtherPerson {
	OtherPerson := []model.OtherPerson{}
	r.DB.Find(&OtherPerson)

	return OtherPerson
}

func (r *repository) GetOtherByID(id int) (Other model.OtherPerson, err error) {
	res := r.DB.Where("id = ?", id).Find(&Other)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateOtherByID(id int, Other model.OtherPerson) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&Other)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteOtherByID(id int) error {
	res := r.DB.Delete(&model.OtherPerson{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}


func NewOtherRepository(db *gorm.DB) domain.AdapterRepositoryOther {
	return &repository{
		DB: db,
	}
}
