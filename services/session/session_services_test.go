package services

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, session model.Session) error}

func (r *repoMockTraditional) CreateSession(session model.Session) error {panic("implement")}
func (r *repoMockTraditional) UpdateSessionByID(hospital_id, vaccine_id int, session model.Session) error{return r.f(hospital_id, session)}
func (r *repoMockTraditional) GetAllSession(hospital_id int) (session []model.Session, err error) {panic("implement")}
func (r *repoMockTraditional) GetSessionByHospitalVaccine(hospital_id, vaccine_id int) (session model.Session, err error) {panic("implement")}
func (r *repoMockTraditional) DeleteSessionByID(hospital_id, vaccine_id int) error {panic("implement")}


func TestUpdateStok(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, stok model.Session) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, stok model.Session) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, stok model.Session) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, stok model.Session) error {
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
	
			svc := NewServiceSession(&repo, config.Config{})
			err := svc.UpdateSessionService(v.id, v.token, model.Session{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}