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

func TestGetAllNews(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewNewsRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `news`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "title", "context"}).
	AddRow(1, "Vaccine Sinovac", "Paling bersahabat").
	AddRow(2, "Vaccine AstraZeneca", "Si Paling kena Rumor").
	AddRow(3, "Covid", "Suka matermatika karena ada alpha, beta, omicron dkk "))

	res := repo.GetAllNews()

	assert.Equal(t, res[0].Title ,"Vaccine Sinovac")
	assert.Len(t, res, 3)
}

func TestGetNewsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewNewsRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `news`")).
	WillReturnRows(sqlmock.NewRows([]string{"id", "title", "context"}).
	AddRow(1, "Vaccine Sinovac", "Paling bersahabat").
	AddRow(2, "Vaccine AstraZeneca", "Si Paling kena Rumor").
	AddRow(3, "Covid", "Suka matermatika karena ada alpha, beta, omicron dkk "))


	res, _ := repo.GetNewsByID(1)

	assert.Equal(t, res, model.News(model.News{ID:0, Title:"", Context:""}))
}

func TestUpdateNewsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewNewsRepository(db)
	defer dbMock.Close()


	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
	WithArgs("abc", 1).
	WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateNewsByID(1, model.News{
		Title: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteNewsByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
			Conn: dbMock,
			SkipInitializeWithVersion: true,
		},
	})
	repo := NewNewsRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
	WithArgs(1).
	WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteNewsByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}