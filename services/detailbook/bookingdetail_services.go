package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcBookingDetail struct {
	c    config.Config
	repo domain.AdapterRepositoryBookingDetail
}

func (s *svcBookingDetail) CreateBookingDetailService(booking model.BookingDetail) error {
	return s.repo.CreateBookingDetail(booking)
}

func (s *svcBookingDetail) GetAllDetailService() []model.BookingDetail {
	return s.repo.GetAllBookingDetail()
}

func (s *svcBookingDetail) GetBookingsByIDService(id int) (booking model.BookingDetail, err error)  {
	return s.repo.GetDetailByID(id)
}

func (s *svcBookingDetail) GetBookingDetailByUserService(user_id int) (booking []model.BookingDetail, err error) {
	return s.repo.GetDetailByUsers(user_id)
}

func (s *svcBookingDetail) GetBookingDetailByOtherPersonService(otherperson_id int) (booking []model.BookingDetail, err error) {
	return s.repo.GetDetailByOtherPerson(otherperson_id)
}

func (s *svcBookingDetail) GetBookingDetailByBookingsService(booking_id int) (booking []model.BookingDetail, err error) {
	return s.repo.GetDetailByBookings(booking_id)
}

func NewServiceBookingDetail(repo domain.AdapterRepositoryBookingDetail, c config.Config) domain.AdapterServiceBookingDetail {
	return &svcBookingDetail{
		repo: repo,
		c:    c,
	}
}
