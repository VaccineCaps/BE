package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, user model.User) error}

func (r *repoMockTraditional) CreateUsers(user model.User) error {panic("implement")}
func (r *repoMockTraditional) GetAll() []model.User {panic("implement")}
func (r *repoMockTraditional) GetOneByID(id int) (user model.User, err error) {panic("implement")}
func (r *repoMockTraditional) GetOneByEmail(email string) (user model.User, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateOneByID(id int, user model.User) error {return r.f(id, user)}
func (r *repoMockTraditional) DeleteByID(id int) error {panic("implement")}
func (r *repoMockTraditional) Login(email, password string) (credential model.User, err error){panic("implement")}

func TestUpdate(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, user model.User) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, user model.User) error {
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
	
			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateUserService(v.id,v.token, model.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestGetOne(t *testing.T) {
// 	testTable := []struct{
// 		text string
// 		f func(id int) (user model.User, err error)
// 		noError bool
// 		token, id int
// 	}{
// 		{
// 			text: "success",
// 			f: func(id int)( user model.User, err error) {return model.User, nil},
// 			noError: true,
// 			token: 1,
// 			id: 1,
// 		},
// 		{
// 			text: "error token != id",
// 			f: func(id int, user model.User) error {
// 				return nil
// 			},
// 			noError: false,
// 			token: 1,
// 			id: 2,
// 		},
// 		{
// 			text: "error internal",
// 			f: func(id int, user model.User) error {
// 				if id == 1 {

// 				}
// 				return errors.New("error")
// 			},
// 			noError: false,
// 			token: 1,
// 			id: 1,
// 		},
// 	}
// 	repo := repoMockTraditional{
// 	}
	
// 	for _, v := range testTable {
// 		t.Run(v.text, func(t *testing.T) {
// 			repo.f = v.f
	
// 			svc := NewServiceUser(&repo, config.Config{})
// 			err := svc.UpdateUserService(v.id,v.token, model.User{})
// 			if v.noError {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }