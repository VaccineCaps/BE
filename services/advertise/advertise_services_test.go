package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, advertise model.Advertise) error}

func (r *repoMockTraditional) CreateAdvertise(advertise model.Advertise) error{panic("implement")}
func (r *repoMockTraditional) GetAllAdvertise() []model.Advertise {panic("implement")}
func (r *repoMockTraditional) GetAdvertiseByID(id int) (advertise model.Advertise, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateAdvertiseByID(id int, advertise model.Advertise) error {return r.f(id, advertise)}
func (r *repoMockTraditional) DeleteAdvertiseByID(id int) error {panic("implement")}

func TestUpdateAdvertise(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, advertise model.Advertise) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, advertise model.Advertise) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, advertise model.Advertise) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, advertise model.Advertise) error {
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
	
			svc := NewServiceAdvertise(&repo, config.Config{})
			err := svc.UpdateAdvertiseService(v.id, model.Advertise{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}