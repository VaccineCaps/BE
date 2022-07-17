package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerVaccineTransaction struct {
	Svc domain.AdapterServiceVaccineTransactionIn
}

func (ce *EchoControllerVaccineTransaction) CreateTransactionHandler(c echo.Context) error {
	transaction := model.VaccineTransactionsIn{}
	c.Bind(&transaction)

	err := ce.Svc.CreateVaccinesTransactionService(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success add transaction",
	})
}

func (ce *EchoControllerVaccineTransaction) UpdateVaccineTransactionController(c echo.Context) error {
	transaction := model.VaccineTransactionsIn{}
	c.Bind(&transaction)

	err := ce.Svc.UpdateTransactionByIDService(transaction.ID, transaction)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccineTransaction) GetTransactionByIDController(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.Svc.GetTransactionByIDService(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":     "success",
		"transactions": res,
	})
}

func (ce *EchoControllerVaccineTransaction) GetAllTransactionInController(c echo.Context) error {
	transaction_in := ce.Svc.GetAllTransactionService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":     "success",
		"transactions": transaction_in,
	}, "  ")
}

func (ce *EchoControllerVaccineTransaction) DeleteVaccineTransactionIDController(c echo.Context) error {
	transaction := model.VaccineTransactionsIn{}
	c.Bind(&transaction)

	err := ce.Svc.DeleteVaccineTransactionByIDService(transaction.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
