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

func (m *MockSvc) CreateRoleService(role model.Role) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) GetAllRoleService() []model.Role {return []model.Role{}}
func (m *MockSvc) GetRoleByID(id int) (model.Role, error) {return model.Role{}, nil}
func (m *MockSvc) DeleteRoleByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateRoleService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateRoleService",  mock.Anything).Return(nil).Once()
	control := EchoControllerRole{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateRoleController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateRoleController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteRoleByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteRoleByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerRole{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteRoleIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteRoleIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}