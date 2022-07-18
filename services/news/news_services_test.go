package service

import (
	"errors"
	"testing"

	"BE/configs"
	"BE/model"
	"github.com/stretchr/testify/assert"
)

type repoMockTraditional struct {f func(id int, news model.News) error}

func (r *repoMockTraditional) CreateNews(news model.News) error{panic("implement")}
func (r *repoMockTraditional) GetAllNews() []model.News {panic("implement")}
func (r *repoMockTraditional) GetNewsByID(id int) (news model.News, err error) {panic("implement")}
func (r *repoMockTraditional) UpdateNewsByID(id int, news model.News) error {return r.f(id, news)}
func (r *repoMockTraditional) DeleteNewsByID(id int) error {panic("implement")}

func TestUpdateNews(t *testing.T) {
	testTable := []struct{
		text string
		f func(id int, news model.News) error
		noError bool
		token, id int
	}{
		{
			text: "success",
			f: func(id int, news model.News) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			text: "error token != id",
			f: func(id int, news model.News) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			text: "error internal",
			f: func(id int, news model.News) error {
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
	
			svc := NewServiceNews(&repo, config.Config{})
			err := svc.UpdateNewsService(v.id, model.News{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}