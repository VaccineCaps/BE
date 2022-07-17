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

func (r *repository) CreateCertificate(vStatus model.Certificate) error {
	res := r.DB.Create(&vStatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllCertificate() []model.Certificate {
	allCertificate := []model.Certificate{}
	r.DB.Find(&allCertificate)

	return allCertificate
}

func (r *repository) GetCertificateByID(id int) (vStatus model.Certificate, err error) {
	res := r.DB.Where("id = ?", id).Find(&vStatus)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateCertificateByID(id int, vStatus model.Certificate) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vStatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}


func (r *repository) DeleteCertificateByID(id int) error {
	res := r.DB.Delete(&model.Certificate{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewCertificateRepository(db *gorm.DB) domain.AdapterRepositoryCertificate {
	return &repository{
		DB: db,
	}
}
