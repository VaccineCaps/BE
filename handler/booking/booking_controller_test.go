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

func (m *MockSvc) CreateBookingService(booking model.Booking) error{
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) GetAllBookingService() []model.Booking{return []model.Booking{}}
func (m *MockSvc) GetBookingByUserService(user_id int) (booking []model.Booking, err error){return booking, nil}
func (m *MockSvc) GetAllBookingBySessionsService(session_id int) (booking []model.Booking, err error) {return booking, nil}
func (m *MockSvc) GetBookingsByIDService(id int) (booking model.Booking, err error) {return booking, nil}
func (m *MockSvc) DeleteBookingByIDService(user_id, hospital_id, session_id, vaccinestatus_id int) error{
	ret := m.Called()
	return ret.Error(0)
}


func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateBookingService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateBookingService",  mock.Anything).Return(nil).Once()
	control := EchoControllerBooking{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateBookingHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateBookingHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteBookingByIDService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteBookingByIDService",  mock.Anything).Return(nil).Once()
	control:= EchoControllerBooking{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteBookingController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteBookingController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}