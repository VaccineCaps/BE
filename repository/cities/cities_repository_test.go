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

func TestGetAllCities(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCityRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `cities`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "province_id"}).
	AddRow(1, "Bandung", 1).
	AddRow(2, "Jakarta Pusat", 1).
	AddRow(3, "Bogor", 1))

	res := repo.GetAllCity()

	assert.Equal(t, res[0].Name ,"Bandung")
	assert.Len(t, res, 3)
}

func TestGetCitiesByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCityRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `cities`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "province_id"}).
	AddRow(1, "Bandung", 1).
	AddRow(2, "Jakarta Pusat", 1).
	AddRow(3, "Bogor", 1))

	res, _ := repo.GetCityByID(1)

	assert.Equal(t, res, model.Cities(model.Cities{ID:0, Name:"", ProvincesId:0 }))
}

func TestDeleteAdvertiseByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCityRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteCityByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}