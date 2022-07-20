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

func (m *MockSvc) CreateCityService(cities model.Cities) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) GetAllCityService() []model.Cities {return []model.Cities{}}
func (m *MockSvc) GetCityByID(id int) (model.Cities, error) {return model.Cities{}, nil}
func (m *MockSvc) DeleteCityByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateCityService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateCityService",  mock.Anything).Return(nil).Once()
	control := EchoControllerCity{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateCityController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateCityController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteCityByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteCityByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerCity{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteCityIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteCityIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}