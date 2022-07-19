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

func (m *MockSvc) CreateSessionService(session model.Session) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateSessionService(hospital_id, vaccine_id int, session model.Session) error{return nil}
func (m *MockSvc) GetAllSessionByHospitalService(hospital_id int) (stok []model.Session, err error) {return []model.Session{}, nil}
func (m *MockSvc) GetSessionByHospitalVaccineService(hospital_id, vaccine_id int) (session model.Session, err error) {return model.Session{}, nil}
func (m *MockSvc) DeleteSessionByIDService(hospital_id, vaccine_id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateSessionService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateSessionService",  mock.Anything).Return(nil).Once()
	control := EchoControllerSession{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateSessionHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateSessionHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteSessionByIDService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteSessionByIDService",  mock.Anything).Return(nil).Once()
	control:= EchoControllerSession{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteSessionIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteSessionIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}