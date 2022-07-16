package repository

import (
	"regexp"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllDetail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewBookingDetailRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookingdetail`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "booking_id", "otherperson_id"}).
	AddRow(1, 1, 1, 1).
	AddRow(2, 1, 2, 1).
	AddRow(3, 2, 1, 1))

	res := repo.GetAllBookingDetail()

	assert.Equal(t, res[0].ID ,1)
	assert.Len(t, res, 3)
}

