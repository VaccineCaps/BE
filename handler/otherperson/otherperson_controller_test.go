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

func (m *MockSvc) CreateOtherService(other model.OtherPerson) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateOtherService(id int, other model.OtherPerson) error {return nil}
func (m *MockSvc) GetAllOtherService() []model.OtherPerson {return []model.OtherPerson{}}
func (m *MockSvc) GetOtherByID(id int) (model.OtherPerson, error) {return model.OtherPerson{}, nil}
func (m *MockSvc) DeleteOtherByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateOtherService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateOtherService",  mock.Anything).Return(nil).Once()
	control := EchoControllerOther{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateOtherController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateOtherController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteOtherByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteOtherByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerOther{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteOtherController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteOtherController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}