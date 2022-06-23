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

func (r *repository) CreateBooking(booking model.Booking) error {

	res := r.DB.Create(&booking)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetAllBooking() []model.Booking {
	allBook := []model.Booking{}
	r.DB.Find(&allBook)

	return allBook
}

func (r *repository) GetAllBookingByUsers(user_id int) (booking []model.Booking, err error) {
	res := r.DB.Where("user_id = ? ", user_id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetAllBookingBySessions(session_id int) (booking []model.Booking, err error) {
	res := r.DB.Where("session_id = ? ", session_id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) GetBookingsByID(id int) (booking model.Booking, err error) {
	res := r.DB.Where("id = ?", id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repository) DeleteBookingByID(user_id, hospital_id, session_id, vaccinestatus_id int) error {
	res := r.DB.Delete(&model.Booking{
		UserId: user_id,
		HospitalId: hospital_id,
		SessionId: session_id,
		VaccineStatusId:vaccinestatus_id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}
func NewBookingRepository(db *gorm.DB) domain.AdapterRepositoryBooking {
	return &repository{
		DB: db,
	}
}
