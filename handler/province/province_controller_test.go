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

func (m *MockSvc) CreateProvinceService(province model.Provinces) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) GetAllProvinceService() []model.Provinces {return []model.Provinces{}}
func (m *MockSvc) GetProvinceByID(id int) (model.Provinces, error) {return model.Provinces{}, nil}
func (m *MockSvc) DeleteProvinceByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateProvinceService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateProvinceService",  mock.Anything).Return(nil).Once()
	control := EchoControllerProvince{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateProvinceController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateProvinceController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteProvinceByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteProvinceByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerProvince{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteProvinceIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteProvinceIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}