package repository

import (
	"fmt"

	"BE/model"
	"gorm.io/gorm"
)

type repositoryDatabaseUser struct {
	DB *gorm.DB
}

func (r *repositoryDatabaseUser) CreateUsers(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryDatabaseUser) GetAll() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *repositoryDatabaseUser) GetOneByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryDatabaseUser) GetOneByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryDatabaseUser) UpdateOneByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repositoryDatabaseUser) DeleteByID(id int) error {
	res := r.DB.Delete(&model.User{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepository(db *gorm.DB) {
	return &repositoryDatabaseUser{
		DB: db,
	}
}