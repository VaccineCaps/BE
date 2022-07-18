package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, vaccine model.Vaccines) error}
type repoMockDelete struct {f func(id int) error}

func (r *repoMockTraditional) CreateVaccines(vaccine model.Vaccines) error{panic("implement")}
func (r *repoMockTraditional) GetAllVaccine() []model.Vaccines {panic("implement")}
func (r *repoMockTraditional) GetVaccineByID(id int) (vaccine model.Vaccines, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateVaccineByID(id int, vaccine model.Vaccines) error {return r.f(id, vaccine)}
func (r *repoMockTraditional) DeleteVaccineByID(id int) error {panic("implement")}

func (r *repoMockDelete) CreateVaccines(vaccine model.Vaccines) error{panic("implement")}
func (r *repoMockDelete) GetAllVaccine() []model.Vaccines {panic("implement")}
func (r *repoMockDelete) GetVaccineByID(id int) (vaccine model.Vaccines, err error) {panic("implement")}
func (r *repoMockDelete) UpdateVaccineByID(id int, vaccine model.Vaccines) error {panic("implement")}
func (r *repoMockDelete) DeleteVaccineByID(id int) error {return r.f(id)}

func TestUpdateVaccines(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, vaccine model.Vaccines) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, vaccine model.Vaccines) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, vaccine model.Vaccines) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, vaccine model.Vaccines) error {
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
	
			svc := NewServiceVaccine(&repo, config.Config{})
			err := svc.UpdateVaccineByID(v.id, model.Vaccines{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

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

	repo := repoMockDelete{}
	
	for _, v := range testTable {
		t.Run(v.text, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceVaccine(&repo, config.Config{})
			err := svc.DeleteVaccineByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}