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

func TestGetAllVaccine(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vaccines`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "Sinovac").
	AddRow(2, "AstraZeneca").
	AddRow(3, "Pfizer"))

	res := repo.GetAllVaccine()

	assert.Equal(t, res[0].Name ,"Sinovac")
	assert.Len(t, res, 3)
}

func TestGetVaccineByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `vaccines`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "Sinovac").
	AddRow(2, "AstraZeneca").
	AddRow(3, "Pfizer"))

	res, _ := repo.GetVaccineByID(1)

	assert.Equal(t, res, model.Vaccines(model.Vaccines{ID:0, Name:""}))
}

func TestUpdateAdvertiseByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateVaccineByID(1, model.Vaccines{
		Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteVaccineByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteVaccineByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}