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

func (m *MockSvc) CreateAdvertiseService(advertise model.Advertise) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateAdvertiseService(id int, advertise model.Advertise) error {return nil}
func (m *MockSvc) GetAllAdvertiseService() []model.Advertise {return []model.Advertise{}}
func (m *MockSvc) GetAdvertiseByID(id int) (model.Advertise, error) {return model.Advertise{}, nil}
func (m *MockSvc) DeleteAdvertiseByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateAdvertiseService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateAdvertiseService",  mock.Anything).Return(nil).Once()
	control := EchoControllerAdvertise{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateAdvertiseController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateAdvertiseController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteAdvertiseByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteAdvertiseByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerAdvertise{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteAdvertiseController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteAdvertiseController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}