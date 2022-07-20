package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, Other model.OtherPerson) error}

func (r *repoMockTraditional) CreateOtherPerson(Other model.OtherPerson) error{panic("implement")}
func (r *repoMockTraditional) GetAllOtherPerson() []model.OtherPerson {panic("implement")}
func (r *repoMockTraditional) GetOtherByID(id int) (Other model.OtherPerson, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateOtherByID(id int, Other model.OtherPerson) error {return r.f(id, Other)}
func (r *repoMockTraditional) DeleteOtherByID(id int) error {panic("implement")}

func TestUpdateCertificate(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, Other model.OtherPerson) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, Other model.OtherPerson) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, Other model.OtherPerson) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, Other model.OtherPerson) error {
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
	
			svc := NewServiceOtherPerson(&repo, config.Config{})
			err := svc.UpdateOtherService(v.id, model.OtherPerson{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}