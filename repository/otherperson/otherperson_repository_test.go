package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"BE/model"
)

func TestGetAllOPs(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewOtherRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `otherperson`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "placeofbirth", "dateofbirth", "address",
	"phonenumber", "nik", "email", "vaccinestatus", "user_id"}).
	AddRow(1, "Azka", "Bandung", "2000-07-13", "Jl. Butuh Pacar No 07",
	"081222009876", 3273000099998888, "ava@abc.com", 0, 1).
	AddRow(2, "Dieva", "Bandung", "1999-07-13", "Jl. Butuh Pacar No 08",
	"081222009876", 3273000099998888, "eva@abc.com", 0, 1).
	AddRow(3, "Via", "Bandung", "2001-07-13", "Jl. Butuh Pacar No 09",
	"081222009876", 3273000099998888, "iva@abc.com", 0, 1))

	res := repo.GetAllOtherPerson()

	assert.Equal(t, res[0].Name ,"Azka")
	assert.Len(t, res, 3)
}

func TestGetOPsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewOtherRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `otherperson`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "placeofbirth", "dateofbirth", "address",
	"phonenumber", "nik", "email", "vaccinestatus", "user_id"}).
	AddRow(1, "Azka", "Bandung", "2000-07-13", "Jl. Butuh Pacar No 07",
	"081222009876", 3273000099998888, "ava@abc.com", 0, 1).
	AddRow(2, "Dieva", "Bandung", "1999-07-13", "Jl. Butuh Pacar No 08",
	"081222009876", 3273000099998888, "eva@abc.com", 0, 1).
	AddRow(3, "Via", "Bandung", "2001-07-13", "Jl. Butuh Pacar No 09",
	"081222009876", 3273000099998888, "iva@abc.com", 0, 1))

	res, _ := repo.GetOtherByID(1)

	assert.Equal(t, res, model.OtherPerson(model.OtherPerson{ID:0, Name:"", Placeofbirth:"", Dateofbirth: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	Address:"", Phonenumber:"", NIK:0, Email:"", VaccineStatus:0, UserId:0 }))
}

func TestUpdateOPsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewOtherRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOtherByID(1, model.OtherPerson{
		Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteOPsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewOtherRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteOtherByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}