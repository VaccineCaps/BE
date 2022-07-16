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

func TestGetAllProvince(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewProvinceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "Jawa Barat").
	AddRow(2, "DKI Jakarta").
	AddRow(3, "Jawa Tengah"))

	res := repo.GetAllProvince()

	assert.Equal(t, res[0].Name, "Jawa Barat")
	assert.Len(t, res, 3)
}

func TestGetProvinceByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewProvinceRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM `provinces`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
	AddRow(1, "Jawa Barat").
	AddRow(2, "DKI Jakarta").
	AddRow(3, "Jawa Tengah"))

	res, _ := repo.GetProvinceByID(1)

	assert.Equal(t, res, model.Provinces(model.Provinces{ID:0, Name:""}))
}


func TestDeleteProvinceByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewProvinceRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteProvinceByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}