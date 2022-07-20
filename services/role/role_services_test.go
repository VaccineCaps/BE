package service

import (
	// "errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int) error}

func (r *repoMockTraditional) CreateRoles(role model.Role) error{panic("implement")}
func (r *repoMockTraditional) GetAllRole() []model.Role {panic("implement")}
func (r *repoMockTraditional) GetRoleByID(id int) (role model.Role, err error) {panic("implement")}
// func (r *repoMockTraditional) UpdateVaccineByID(id int, role model.Role) error {return r.f(id, role)}
func (r *repoMockTraditional) DeleteRoleByID(id int) error {return r.f(id)}

func TestDeleteRole(t *testing.T) {
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
	
			svc := NewServiceRole(&repo, config.Config{})
			err := svc.DeleteRoleByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}