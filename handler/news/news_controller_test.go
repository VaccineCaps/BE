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

func (m *MockSvc) CreateNewsService(news model.News) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateNewsService(id int, news model.News) error {return nil}
func (m *MockSvc) GetAllNewsService() []model.News {return []model.News{}}
func (m *MockSvc) GetNewsByID(id int) (model.News, error) {return model.News{}, nil}
func (m *MockSvc) DeleteNewsByID(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateNewsService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateNewsService",  mock.Anything).Return(nil).Once()
	control := EchoControllerNews{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateNewsController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateNewsController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteNewsByID",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteNewsByID",  mock.Anything).Return(nil).Once()
	control:= EchoControllerNews{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteNewsController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteNewsController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}