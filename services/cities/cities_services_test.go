package service

import (
	// "errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int) error}

func (r *repoMockTraditional) CreateCities(cities model.Cities) error{panic("implement")}
func (r *repoMockTraditional) GetAllCity() []model.Cities {panic("implement")}
func (r *repoMockTraditional) GetCityByID(id int) (cities model.Cities, err error) {panic("implement")}
// func (r *repoMockTraditional) UpdateVaccineByID(id int, cities model.Cities) error {return r.f(id, role)}
func (r *repoMockTraditional) DeleteCityByID(id int) error {return r.f(id)}

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
	
			svc := NewServiceCity(&repo, config.Config{})
			err := svc.DeleteCityByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}