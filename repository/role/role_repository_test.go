package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"BE/model"
)

func TestGetAllRole(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewRoleRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `role`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "admin").
	AddRow(2, "user"))

	res := repo.GetAllRole()

	assert.Equal(t, res[0].Name, "admin")
	assert.Len(t, res, 2)
}

func TestGetRoleByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewRoleRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM `role`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "admin").
	AddRow(2, "user"))

	res, _ := repo.GetRoleByID(1)

	assert.Equal(t, res, model.Role(model.Role{ID:0, Name:""}))
}


func TestDeleteRoleByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewRoleRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteRoleByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}