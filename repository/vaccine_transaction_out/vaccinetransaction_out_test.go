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

func TestGetAllOut(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineTransactionRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vaccine_transactions_out`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccinehospital_id", "status", "tanggal", "no_transaction", "distributor"}).
	AddRow(1, 1, 1, 1, "2022-07-13", 1,"Kimia Farma" ).
	AddRow(2, 1, 1, 0, "2022-07-26", 2,"Kimia Farma").
	AddRow(3, 1, 1, 0, "2022-07-13", 3,"Kimia Farma"))

	res := repo.GetAllTransaction()

	assert.Equal(t, res[0].Distributor,"")
	assert.Len(t, res, 3)
}

func TestGetOutByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineTransactionRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `vaccine_transactions_out`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccinehospital_id", "status", "tanggal", "no_transaction", "distributor"}).
	AddRow(1, 1, 1, 1, "2022-07-13", 1,"Kimia Farma" ).
	AddRow(2, 1, 1, 0, "2022-07-26", 2,"Kimia Farma").
	AddRow(3, 1, 1, 0, "2022-07-13", 3,"Kimia Farma"))

	res, _ := repo.GetTransactionByID(1)

	assert.Equal(t, res, []model.VaccineTransactionsOut([]model.VaccineTransactionsOut(nil)))
}

func TestUpdateOutByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineTransactionRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateTransactionByID(1, model.VaccineTransactionsOut{
		Distributor: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteOutByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineTransactionRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteVaccineTransactionByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}