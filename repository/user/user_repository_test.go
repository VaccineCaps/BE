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

func TestGetAll(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password","image","RoleId"}).
	AddRow(1, "Azifaazka", "zifa@abc.com", "12344321", "png", 1).
	AddRow(2, "irfancahyo", "irfan@abc.com", "12344321", "png", 1).
	AddRow(3, "alfaisal", "faisal@abc.com", "12344321", "png", 2))

	res := repo.GetAll()

	assert.Equal(t, res[0].Username ,"Azifaazka")
	assert.Len(t, res, 3)
}

func TestGetByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `users`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password","image","RoleId"}).
	AddRow(1, "Azifaazka", "zifa@abc.com", "12344321", "png", 1).
	AddRow(2, "irfancahyo", "irfan@abc.com", "12344321", "png", 1).
	AddRow(3, "alfaisal", "faisal@abc.com", "12344321", "png", 2))


	res, _ := repo.GetOneByID(1)

	assert.Equal(t, res, model.User(model.User{ID:0, Username:"", Email:"", Password:"", Image:"", RoleId:0}))
}

func TestUpdateByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewUserRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneByID(1, model.User{
		Username: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteNewsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}