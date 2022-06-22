package repository

import (
	"BE/domain"
	"BE/model"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateNews(news model.News) error {

	res := r.DB.Create(&news)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllNews() []model.News {
	news := []model.News{}
	r.DB.Find(&news)

	return news
}

func (r *repository) GetNewsByID(id int) (news model.News, err error) {
	res := r.DB.Where("id = ?", id).Find(&news)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateNewsByID(id int, news model.News) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&news)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteNewsByID(id int) error {
	res := r.DB.Delete(&model.News{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewNewsRepository(db *gorm.DB) domain.AdapterRepositoryNews {
	return &repository{
		DB: db,
	}
}
