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

func (m *MockSvc) CreateVaccinesTransactionService(transaction model.VaccineTransactionsIn) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateTransactionByIDService(id int, transaction model.VaccineTransactionsIn) error {return nil}
func (m *MockSvc) GetAllTransactionService() []model.VaccineTransactionsIn {return []model.VaccineTransactionsIn{}}
func (m *MockSvc) GetTransactionByIDService(id int) (transaction []model.VaccineTransactionsIn, err error) {return transaction, nil}
func (m *MockSvc) DeleteVaccineTransactionByIDService(id int) error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateController(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateVaccinesTransactionService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateVaccinesTransactionService",  mock.Anything).Return(nil).Once()
	control := EchoControllerVaccineTransaction{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		control.CreateTransactionHandler(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		
		control.CreateTransactionHandler(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteVaccineTransactionByIDService",  mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteVaccineTransactionByIDService",  mock.Anything).Return(nil).Once()
	control:= EchoControllerVaccineTransaction{Svc: &svc}

	e := echo.New()
	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteVaccineTransactionIDController(echoContext)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		control.DeleteVaccineTransactionIDController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}