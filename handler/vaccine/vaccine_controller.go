package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerVaccine struct {
	Svc domain.AdapterServiceVaccine
}

func (ce *EchoControllerVaccine) CreateVaccineController(c echo.Context) error {
	vaccine := model.Vaccines{}
	c.Bind(&vaccine)

	err := ce.Svc.CreateVaccineService(vaccine)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
	})
}

func (ce *EchoControllerVaccine) GetVaccineIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetVaccineByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success",
		"provinces": res,
	})
}

func (ce *EchoControllerVaccine) GetAllVaccineController(c echo.Context) error {
	vaccines := ce.Svc.GetAllVaccineService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":  "success",
		"provinces": vaccines,
	}, "  ")
}

func (ce *EchoControllerVaccine) UpdateVaccineController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vaccine := model.Vaccines{}
	c.Bind(&vaccine)

	err := ce.Svc.UpdateVaccineByID(intID, vaccine)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccine) DeleteVaccineIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteVaccineByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
