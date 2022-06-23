package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcBooking struct {
	c    config.Config
	repo domain.AdapterRepositoryBooking
}

func (s *svcBooking) CreateBookingService(booking model.Booking) error {
	return s.repo.CreateBooking(booking)
}

func (s *svcBooking) GetAllBookingService() []model.Booking {
	return s.repo.GetAllBooking()
}

func (s *svcBooking) GetBookingByUserService(user_id int) (booking []model.Booking, err error) {
	return s.repo.GetAllBookingByUsers(user_id)
}

func (s *svcBooking) GetAllBookingBySessionsService(session_id int) (booking []model.Booking, err error) {
	return s.repo.GetAllBookingBySessions(session_id)
}

func (s *svcBooking) GetBookingsByIDService(id int) (booking model.Booking, err error)  {
	return s.repo.GetBookingsByID(id)
}

func (s *svcBooking) DeleteBookingByIDService(user_id, hospital_id, session_id, vaccinestatus_id int) error {
	return s.repo.DeleteBookingByID(user_id, hospital_id, session_id, vaccinestatus_id)
}

func NewServiceBooking(repo domain.AdapterRepositoryBooking, c config.Config) domain.AdapterServiceBooking {
	return &svcBooking{
		repo: repo,
		c:    c,
	}
}
