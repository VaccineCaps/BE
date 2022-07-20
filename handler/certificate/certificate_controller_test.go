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

func (m *MockSvc) CreateCertificateService(sertif model.Certificate) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateCertificateService(id int, sertif model.Certificate) error {return nil}
func (m *MockSvc) GetAllCertificateService() []model.Certificate {return []model.Certificate{}}
func (m *MockSvc) GetCertificateByID(id int) (model.Certificate, error) {return model.Certificate{}, nil}
func (m *MockSvc) DeleteCertificateByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateCertificateService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateCertificateService",  mock.Anything).Return(nil).Once()
	control := EchoControllerCertificate{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateCertificateController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateCertificateController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteCertificateByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteCertificateByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerCertificate{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteCertificateIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteCertificateIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}