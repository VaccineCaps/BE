package services

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, transaction model.VaccineTransactionsOut) error}

func (r *repoMockTraditional) CreateVaccinesTransaction(transaction model.VaccineTransactionsOut) error{panic("implement")}
func (r *repoMockTraditional) GetAllTransaction() []model.VaccineTransactionsOut {panic("implement")}
func (r *repoMockTraditional) GetTransactionByID(id int) (transaction []model.VaccineTransactionsOut, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateTransactionByID(id int, transaction model.VaccineTransactionsOut) error {return r.f(id, transaction)}
func (r *repoMockTraditional) DeleteVaccineTransactionByID(id int) error {panic("implement")}

func TestUpdateVaccines(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, transaction model.VaccineTransactionsOut) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, transaction model.VaccineTransactionsOut) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, transaction model.VaccineTransactionsOut) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, transaction model.VaccineTransactionsOut) error {
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
	
			svc := NewServiceVaccineTransactionsOut(&repo, config.Config{})
			err := svc.UpdateTransactionByIDService(v.id, model.VaccineTransactionsOut{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

type repoMockDelete struct {f func(id int) error}
func (r *repoMockDelete) CreateVaccinesTransaction(transaction model.VaccineTransactionsOut) error{panic("implement")}
func (r *repoMockDelete) GetAllTransaction() []model.VaccineTransactionsOut {panic("implement")}
func (r *repoMockDelete) GetTransactionByID(id int) (transaction []model.VaccineTransactionsOut, err error) {panic("implement")}
func (r *repoMockDelete) UpdateTransactionByID(id int, transaction model.VaccineTransactionsOut) error {panic("implement")}
func (r *repoMockDelete) DeleteVaccineTransactionByID(id int) error {return r.f(id)}

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

	repo := repoMockDelete{}
	
	for _, v := range testTable {
		t.Run(v.text, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceVaccineTransactionsOut(&repo, config.Config{})
			err := svc.DeleteVaccineTransactionByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}