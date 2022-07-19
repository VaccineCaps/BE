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

func (m *MockSvc) CreateHospitalService(hospital model.Hospitals) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateHospitalService(id int, hospital model.Hospitals) error {return nil}
func (m *MockSvc) GetAllHospitalsService() []model.Hospitals {return []model.Hospitals{}}
func (m *MockSvc) GetHospitalsByID(id int) (model.Hospitals, error) {return model.Hospitals{}, nil}
func (m *MockSvc) DeleteHospitalByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateHospitalService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateHospitalService",  mock.Anything).Return(nil).Once()
	control := EchoControllerHospital{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateHospitalController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateHospitalController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteHospitalByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteHospitalByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerHospital{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteHospitalController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteHospitalController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}