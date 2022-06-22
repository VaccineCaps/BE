package repository

import (
	"BE/domain"
	"BE/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateUsers(user model.User) error {

	hashedPassword, _ := hashPassword(user.Password)
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

	match := checkPasswordHash(password, hashedPassword)
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

// hashPassword implement bycrypt for password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash check bycrypted password
func checkPasswordHash(password, hashed string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}

func NewUserRepository(db *gorm.DB) domain.AdapterRepositoryUser {
	return &repository{
		DB: db,
	}
}
