package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerVaccineTransaction struct {
	Svc domain.AdapterServiceVaccineTransaction
}

func (ce *EchoControllerVaccineTransaction) CreateTransactionHandler(c echo.Context) error {
	transaction := model.VaccineTransactions{}
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
	transaction := model.VaccineTransactions{}
	c.Bind(&transaction)

	err := ce.Svc.UpdateTransactionByIDService(transaction.HospitalId, transaction.VaccineId, transaction)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccineTransaction) GetTrnasactionByHospitalController(c echo.Context) error {
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	if err != nil {

	}

	res, err := ce.Svc.GetAllTransactionByHospitalService(HospitalID)
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

func (ce *EchoControllerVaccineTransaction) GetTransactionByHospitalVaccineIDController(c echo.Context) error {
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	vaccine_id := c.Param("vaccine_id")
	VaccineID, err := strconv.Atoi(vaccine_id)
	if err != nil {

	}

	res, err := ce.Svc.GetTransactionByHospitalVaccineService(HospitalID, VaccineID)
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

func (ce *EchoControllerVaccineTransaction) DeleteVaccineTransactionIDController(c echo.Context) error {
	transaction := model.VaccineTransactions{}
	c.Bind(&transaction)

	err := ce.Svc.DeleteVaccineTransactionByIDService(transaction.HospitalId, transaction.VaccineId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
