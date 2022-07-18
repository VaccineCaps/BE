package repository

import (
	// "database/sql/driver"
	"regexp"
	"time"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"BE/model"
)

func TestGetAllSession(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewSessionRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `hospital_id` FROM `session`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccine_id", "sesi", "max_session", "number_booking", "tanggal"}).
	AddRow(1, 1, 1, "Pagi", 50,50, "2022-07-13").
	AddRow(2, 1, 1, "Siang", 50,50, "2022-07-13").
	AddRow(3, 1, 1, "Sore", 50,50, "2022-07-13"))

	res, _ := repo.GetAllSession(1)

	assert.Equal(t, res, []model.Session([]model.Session(nil)))
}

func TestGetSessionByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewSessionRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `hospital_id`, `vaccine_id` FROM `session`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccine_id", "sesi", "max_session", "number_booking", "tanggal"}).
	AddRow(1, 1, 1, "Pagi", 50,50, "2022-07-13").
	AddRow(2, 1, 1, "Siang", 50,50, "2022-07-13").
	AddRow(3, 1, 1, "Sore", 50,50, "2022-07-13"))

	res, _ := repo.GetSessionByHospitalVaccine(1,1)

	assert.Equal(t, res, model.Session(model.Session{ID:0, HospitalId:0, VaccineId:0, Sesi:"", MaxSession:0, NumberBooking:0,
	Tanggal:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}))
}

// func TestUpdateAdvertiseByID(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 			Conn: dbMock,
// 			SkipInitializeWithVersion: true,
// 		},
// 	})
// 	repo := NewSessionRepository(db)
// 	defer dbMock.Close()


// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
// 	WithArgs(50, 1).
// 	WillReturnResult(sqlmock.NewResult(0, 1))
// 	fMock.ExpectCommit()

// 	err := repo.UpdateSessionByID(1, 1, model.Session{
// 		MaxSession: 50,
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

// func TestDeleteVaccineByID(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 			Conn: dbMock,
// 			SkipInitializeWithVersion: true,
// 		},
// 	})
// 	repo := NewSessionRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
// 	WithArgs(1,1).
// 	WillReturnResult(driver.RowsAffected(1))
// 	fMock.ExpectCommit()

// 	err:= repo. DeleteSessionByID(1,1)
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }