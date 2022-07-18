package services

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, stok model.VaccineHospitals) error}

func (r *repoMockTraditional) CreateStokVaccineHospital(stok model.VaccineHospitals) error {panic("implement")}
func (r *repoMockTraditional) UpdateStokByID(hospital_id, vaccine_id int, stok model.VaccineHospitals) error {return r.f(hospital_id, stok)}
func (r *repoMockTraditional) GetAllStokByHospital(hospital_id int) (stok []model.VaccineHospitals, err error) {panic("implement")}
func (r *repoMockTraditional) GetStokByHospitalVaccine(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error) {panic("implement")}
func (r *repoMockTraditional) DeleteVaccineStokByID(hospital_id, vaccine_id int) error {panic("implement")}

func TestUpdateStok(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, stok model.VaccineHospitals) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, stok model.VaccineHospitals) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, stok model.VaccineHospitals) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, stok model.VaccineHospitals) error {
				if id == 1 {

				}
				return errors.New("error")
			},
			noError: false,
			token: 1,
			id: 1,
		},
	}
	repo := repoMockTraditional{
	}
	
	for _, v := range testTable {
		t.Run(v.text, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceVaccineHospital(&repo, config.Config{})
			err := svc.UpdateVaccineStokService(v.id, v.token, model.VaccineHospitals{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}