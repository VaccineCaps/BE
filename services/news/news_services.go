package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcNews struct {
	c    config.Config
	repo domain.AdapterRepositoryNews
}

func (s *svcNews) CreateNewsService(News model.News) error {
	return s.repo.CreateNews(News)
}

func (s *svcNews) UpdateNewsService(id int, News model.News) error {
	return s.repo.UpdateNewsByID(id, News)
}

func (s *svcNews) GetAllNewsService() []model.News {
	return s.repo.GetAllNews()
}

func (s *svcNews) GetNewsByID(id int) (model.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *svcNews) DeleteNewsByID(id int) error {
	return s.repo.DeleteNewsByID(id)
}

func NewServiceNews(repo domain.AdapterRepositoryNews, c config.Config) domain.AdapterServiceNews {
	return &svcNews{
		repo: repo,
		c:    c,
	}
}
