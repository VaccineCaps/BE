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

func (m *MockSvc) CreateVaccineService(vaccine model.Vaccines) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateVaccineByID(id int, vaccine model.Vaccines) error {return nil}
func (m *MockSvc) GetAllVaccineService() []model.Vaccines {return []model.Vaccines{}}
func (m *MockSvc) GetVaccineByID(id int) (model.Vaccines, error) {return model.Vaccines{}, nil}
func (m *MockSvc) DeleteVaccineByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateVaccineService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateVaccineService",  mock.Anything).Return(nil).Once()
	control := EchoControllerVaccine{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateVaccineController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateVaccineController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteVaccineByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteVaccineByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerVaccine{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteVaccineIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteVaccineIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}