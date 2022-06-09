package repository

import (
	"fmt"

	"BE/domain"
	"BE/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateRoles(role model.Role) error {
	res := r.DB.Create(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllRole() []model.Role {
	roles := []model.Role{}
	r.DB.Find(roles)

	return roles
}

func (r *repository) GetRoleByID(id int) (role model.Role, err error) {
	res := r.DB.Where("id = ?", id).Find(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetRoleByName(name string) (role model.Role, err error) {
	res := r.DB.Where("name = ?", name).Find(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteRoleByID(id int) error {
	res := r.DB.Delete(&model.Role{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func (r *repository) DeleteRoleByName(name string) error {
	res := r.DB.Delete(&model.Role{
		Name: name,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewRoleRepository(db *gorm.DB) domain.AdapterRepositoryRole {
	return &repository{
		DB: db,
	}
}
