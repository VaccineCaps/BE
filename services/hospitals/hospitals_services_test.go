package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, hospital model.Hospitals) error}

func (r *repoMockTraditional) CreateHospitals(hospital model.Hospitals) error{panic("implement")}
func (r *repoMockTraditional) GetAllHospitals() []model.Hospitals {panic("implement")}
func (r *repoMockTraditional) GetHospitalByID(id int) (hospital model.Hospitals, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateHospitalByID(id int, hospital model.Hospitals) error {return r.f(id, hospital)}
func (r *repoMockTraditional) DeleteHospitalByID(id int) error {panic("implement")}

func TestUpdateHospital(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, hospital model.Hospitals) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, hospital model.Hospitals) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, hospital model.Hospitals) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, hospital model.Hospitals) error {
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
	
			svc := NewServiceHospitals(&repo, config.Config{})
			err := svc.UpdateHospitalService(v.id, model.Hospitals{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}