package repository

import (
	"BE/domain"
	"BE/model"
	repo "BE/repository/hash"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateUsers(user model.User) error {

	hashedPassword, _ := repo.HashPassword(user.Password)
	user.Password = hashedPassword
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) Login(email, password string) (credential model.User, err error) {
	var hashedPassword string
	_ = r.DB.Raw("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)

	match := repo.CheckPasswordHash(password, hashedPassword)
	if !match {
		return credential, fmt.Errorf("invalid credentials")
	}

	response := r.DB.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&credential)

	if response.RowsAffected < 1 {
		return credential, fmt.Errorf("email or password is incorrect")
	}

	return credential, nil
}

func (r *repository) GetAll() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *repository) GetOneByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetOneByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) UpdateOneByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repository) DeleteByID(id int) error {
	res := r.DB.Delete(&model.User{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewUserRepository(db *gorm.DB) domain.AdapterRepositoryUser {
	return &repository{
		DB: db,
	}
}
