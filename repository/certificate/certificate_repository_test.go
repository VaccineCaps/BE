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

func TestGetAllCertificate(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCertificateRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `certificate`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "image"}).
	AddRow(1, "Vaccine Sinovac").
	AddRow(2, "Vaccine AstraZeneca").
	AddRow(3, "Covid"))

	res := repo.GetAllCertificate()

	assert.Equal(t, res[0].Image ,"Vaccine Sinovac")
	assert.Len(t, res, 3)
}

func TestGetCertificateByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCertificateRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `certificate`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "image"}).
	AddRow(1, "Vaccine Sinovac").
	AddRow(2, "Vaccine AstraZeneca").
	AddRow(3, "Covid"))

	res, _ := repo.GetCertificateByID(1)

	assert.Equal(t, res, model.Certificate(model.Certificate{ID:0, Image:""}))
}

func TestUpdateCertificateByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCertificateRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateCertificateByID(1, model.Certificate{
		Image: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteCertificateByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewCertificateRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteCertificateByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}