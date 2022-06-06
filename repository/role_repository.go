package repository

import (
	"fmt"

	"BE/model"
	"gorm.io/gorm"
)

type repositoryDatabaseRole struct {
	DB *gorm.DB
}

func (r *repositoryDatabaseRole) CreateRoles(role model.Role) error {
	res := r.DB.Create(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryDatabaseRole) GetAllRole() []model.Role {
	roles := []model.Role{}
	r.DB.Find(roles)

	return roles
}

func (r *repositoryDatabaseRole) GetRoleByID(id int) (role model.Role, err error) {
	res := r.DB.Where("id = ?", id).Find(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryDatabaseRole) GetRoleByName(name string) (role model.Role, err error) {
	res := r.DB.Where("name = ?", name).Find(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryDatabaseRole) DeleteRoleByID(id int) error {
	res := r.DB.Delete(&model.Role{
		ID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func (r *repositoryDatabaseRole) DeleteRoleByName(id int) error {
	res := r.DB.Delete(&model.Role{
		Name: name,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryDatabaseRole{
		DB: db,
	}
}