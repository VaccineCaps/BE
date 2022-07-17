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

func TestGetAllBookings(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewBookingRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `booking`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "hospital_id", "session_id", "certificate_id",
	"bookedcode", "statusbook", "vaccinenumber", "statusnumber"}).
	AddRow(1, 1, 1, 1, 1, "0001", 0, 1, 0 ).
	AddRow(2, 1, 1, 1, 1, "0002", 0, 1, 0).
	AddRow(3, 1, 1, 1, 1, "0003", 0, 1, 0))

	res := repo.GetAllBooking()

	assert.Equal(t, res[0].BookedCode ,"0001")
	assert.Len(t, res, 3)
}

func TestGetBookingsyID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewBookingRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `booking`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "hospital_id", "session_id", "certificate_id",
	"bookedcode", "statusbook", "vaccinenumber", "statusnumber"}).
	AddRow(1, 1, 1, 1, 1, "0001", 0, 1, 0 ).
	AddRow(2, 1, 1, 1, 1, "0002", 0, 1, 0).
	AddRow(3, 1, 1, 1, 1, "0003", 0, 1, 0))


	res, _ := repo.GetBookingsByID(1)

	assert.Equal(t, res, model.Booking(model.Booking{ID:0, UserId:0, HospitalId:0, SessionId:0, CertificateId:0, 
	BookedCode:"", StatusBook:false, VaccineNumber:0, StatusNumber:false}))
}

func TestDeleteBookingssByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewBookingRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1,1,1,1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteBookingByID(1,1,1,1)
	assert.NoError(t, err)
	assert.True(t, true)
}