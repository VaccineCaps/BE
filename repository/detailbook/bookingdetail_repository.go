package repository

import (
	"BE/domain"
	"BE/model"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (r *repository) CreateBookingDetail(book model.BookingDetail) error {

	res := r.DB.Create(&book)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetDetailByID(id int) (booking model.BookingDetail, err error) {
	res := r.DB.Where("id = ?", id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetAllBookingDetail() []model.BookingDetail {
	alldetail := []model.BookingDetail{}
	r.DB.Find(&alldetail)

	return alldetail
}

func (r *repository) GetDetailByUsers(user_id int) (booking []model.BookingDetail, err error) {
	res := r.DB.Where("user_id = ? ", user_id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetDetailByOtherPerson(otherperson_id int) (booking []model.BookingDetail, err error) {
	res := r.DB.Where("otherperson_id = ? ", otherperson_id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetDetailByBookings(booking_id int) (booking []model.BookingDetail, err error) {
	res := r.DB.Where("booking_id = ? ", booking_id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func NewBookingDetailRepository(db *gorm.DB) domain.AdapterRepositoryBookingDetail {
	return &repository{
		DB: db,
	}
}
