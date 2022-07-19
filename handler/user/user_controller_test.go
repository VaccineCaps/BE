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
func (m *MockSvc) CreateUserService(user model.User) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateUserService(id, idToken int, user model.User) error {return nil}
func (m *MockSvc) GetAllUsersService() []model.User {return []model.User{}}
func (m *MockSvc) GetUserByID(id int) (model.User, error) {return model.User{}, nil}
func (m *MockSvc) LoginUser(email, password string) (string, int) {return "success", 200}
func (m *MockSvc) DeleteByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestRegisterController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateUserService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateUserService",  mock.Anything).Return(nil).Once()
	usrController := EchoControllerUser{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.RegisterHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.RegisterHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteByID",  mock.Anything).Return(nil).Once()
	usrController := EchoControllerUser{Svc: &svc,}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}