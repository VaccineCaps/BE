package service

import (
	// "errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int) error}

func (r *repoMockTraditional) CreateProvinces(province model.Provinces) error{panic("implement")}
func (r *repoMockTraditional) GetAllProvince() []model.Provinces {panic("implement")}
func (r *repoMockTraditional) GetProvinceByID(id int) (province model.Provinces, err error) {panic("implement")}
// func (r *repoMockTraditional) UpdateVaccineByID(id int, province model.Provinces) error {return r.f(id, role)}
func (r *repoMockTraditional) DeleteProvinceByID(id int) error {return r.f(id)}

func TestDelete(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int) error {
				return nil
			},
			noError: true,
			id: 1,
		},
		{
			text: "id not found",
			f: func(id int) error {
				return nil
			},
			noError: false,
			
		},
	}

	repo := repoMockTraditional{}
	
	for _, v := range testTable {
		t.Run(v.text, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceProvince(&repo, config.Config{})
			err := svc.DeleteProvinceByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}