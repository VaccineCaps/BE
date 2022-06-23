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

func (r *repository) CreateSession(session model.Session) error {

	res := r.DB.Create(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) UpdateSessionByID(hospital_id, vaccine_id int, session model.Session) error {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).UpdateColumns(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) GetAllSession(hospital_id int) (session []model.Session, err error) {
	res := r.DB.Where("hospital_id = ? ", hospital_id).Find(&session)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetSessionByHospitalVaccine(hospital_id, vaccine_id int) (session model.Session, err error) {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).Find(&session)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteSessionByID(hospital_id, vaccine_id int) error {
	res := r.DB.Delete(&model.Session{
		HospitalId: hospital_id,
		VaccineId:  vaccine_id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewSessionRepository(db *gorm.DB) domain.AdapterRepositorySession {
	return &repository{
		DB: db,
	}
}
