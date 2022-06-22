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

func (r *repository) CreateStokVaccineHospital(stok model.VaccineHospitals) error {

	res := r.DB.Create(&stok)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) UpdateStokByID(hospital_id, vaccine_id int, stok model.VaccineHospitals) error {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).UpdateColumns(&stok)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) GetAllStokByHospital(hospital_id int) (stok []model.VaccineHospitals, err error) {
	res := r.DB.Where("hospital_id = ? ", hospital_id).Find(&stok)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetStokByHospitalVaccine(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error) {
	res := r.DB.Where("hospital_id = ? AND vaccine_id = ?", hospital_id, vaccine_id).Find(&stok)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteVaccineStokByID(hospital_id, vaccine_id int) error {
	res := r.DB.Delete(&model.VaccineHospitals{
		HospitalId: hospital_id,
		VaccineId:  vaccine_id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}
func NewVaccineHospitalRepository(db *gorm.DB) domain.AdapterRepositoryVaccineHospital {
	return &repository{
		DB: db,
	}
}
