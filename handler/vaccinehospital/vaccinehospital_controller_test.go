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

func (m *MockSvc) CreateStokService(stok model.VaccineHospitals) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateVaccineStokService(hospital_id, vaccine_id int, stok model.VaccineHospitals) error{return nil}
func (m *MockSvc) GetAllStokByHospitalService(hospital_id int) (stok []model.VaccineHospitals, err error) {return []model.VaccineHospitals{}, nil}
func (m *MockSvc) GetStokByHospitalVaccineService(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error) {return model.VaccineHospitals{}, nil}
func (m *MockSvc) DeleteVaccineStokByIDService(hospital_id, vaccine_id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateStokService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateStokService",  mock.Anything).Return(nil).Once()
	control := EchoControllerVaccineHospital{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateStokHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateStokHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteVaccineStokByIDService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteVaccineStokByIDService",  mock.Anything).Return(nil).Once()
	control:= EchoControllerVaccineHospital{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteHospitalVaccineIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteHospitalVaccineIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}