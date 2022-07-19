package handler

import (
	"errors"
	"net/http/httptest"
	"testing"
	"BE/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockSvc struct {mock.Mock}

func (m *MockSvc) CreateBookingDetailService(booking model.BookingDetail) error{
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) GetAllDetailService() []model.BookingDetail{return []model.BookingDetail{}}
func (m *MockSvc) GetBookingsByIDService(id int) (booking model.BookingDetail, err error){return booking, nil}
func (m *MockSvc) GetBookingDetailByUserService(user_id int) (booking []model.BookingDetail, err error){return booking, nil}
func (m *MockSvc) GetBookingDetailByOtherPersonService(otherperson_id int) (booking []model.BookingDetail, err error){return booking, nil}
func (m *MockSvc) GetBookingDetailByBookingsService(booking_id int) (booking []model.BookingDetail, err error){return booking, nil}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateBookingDetailService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateBookingDetailService",  mock.Anything).Return(nil).Once()
	control := EchoControllerBookingDetail{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateBookingDetailHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateBookingDetailHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}
