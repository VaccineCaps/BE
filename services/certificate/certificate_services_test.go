package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, sertif model.Certificate) error}

func (r *repoMockTraditional) CreateCertificate(sertif model.Certificate) error{panic("implement")}
func (r *repoMockTraditional) GetAllCertificate() []model.Certificate {panic("implement")}
func (r *repoMockTraditional) GetCertificateByID(id int) (sertif model.Certificate, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateCertificateByID(id int, sertif model.Certificate) error {return r.f(id, sertif)}
func (r *repoMockTraditional) DeleteCertificateByID(id int) error {panic("implement")}

func TestUpdateCertificate(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, sertif model.Certificate) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, sertif model.Certificate) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, sertif model.Certificate) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, sertif model.Certificate) error {
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
	
			svc := NewServiceCertificate(&repo, config.Config{})
			err := svc.UpdateCertificateService(v.id, model.Certificate{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}