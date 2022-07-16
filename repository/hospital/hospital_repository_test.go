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

func TestGetAllHospital(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewHospitalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hospitals`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "email", "user_id", "cities_id"}).
	AddRow(1, "RS Martadinata", "Jl. aku padamu no 12", "marta@abc.com", 1, 1).
	AddRow(2, "RS Cinta", "Jl. Status no 12", "status@abc.com", 1, 2).
	AddRow(3, "RS Penyakit Hati", "Jl. diatigakali no 12", "dia@abc.com", 1, 3))

	res := repo.GetAllHospitals()

	assert.Equal(t, res[0].Name ,"RS Martadinata")
	assert.Len(t, res, 3)
}

func TestGetHospitalID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewHospitalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `hospitals`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "email", "user_id", "cities_id"}).
	AddRow(1, "RS Martadinata", "Jl. aku padamu no 12", "marta@abc.com", 1, 1).
	AddRow(2, "RS Cinta", "Jl. Status no 12", "status@abc.com", 1, 2).
	AddRow(3, "RS Penyakit Hati", "Jl. diatigakali no 12", "dia@abc.com", 1, 3))

	res, _ := repo.GetHospitalByID(1)

	assert.Equal(t, res, model.Hospitals(model.Hospitals{ID:0, Name:"", Email:"", UserId:0, CitiesId:0}))
}

func TestUpdateHospitalByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewHospitalRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateHospitalByID(1, model.Hospitals{
		Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteHospitalByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewHospitalRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteHospitalByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}