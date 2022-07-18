package repository

import (
	// "database/sql/driver"
	"regexp"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"BE/model"
)

func TestGetAllStokByIDHospital(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineHospitalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `hospital_id` FROM `vaccine_hospitals`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccine_id", "stok"}).
	AddRow(1, 1,1,50).
	AddRow(2, 1,2,100).
	AddRow(3, 2, 1, 200))

	res, _ := repo.GetAllStokByHospital(1)
	assert.Equal(t, res, []model.VaccineHospitals([]model.VaccineHospitals(nil)))
	
}

func TestGetStokByHospitalVaccine(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewVaccineHospitalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `hospitals_id`, `vaccine_id` FROM `vaccine_hospitals`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "hospital_id", "vaccine_id", "stok"}).
	AddRow(1, 1, 1,50).
	AddRow(2, 1, 2,100).
	AddRow(3, 2, 1, 200))

	res, _ := repo.GetStokByHospitalVaccine(1,1)
	assert.Equal(t, res, model.VaccineHospitals(model.VaccineHospitals{ID:0, HospitalId:0, VaccineId:0, Stok:0}))
}



// func TestDeleteStokByID(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 			Conn: dbMock,
// 			SkipInitializeWithVersion: true,
// 		},
// 	})
// 	repo := NewVaccineHospitalRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
// 	WithArgs(1,1).
// 	WillReturnResult(driver.RowsAffected(1))
// 	fMock.ExpectCommit()

// 	err := repo.DeleteVaccineStokByID(1,1)
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }