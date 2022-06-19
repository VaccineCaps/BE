package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerHospital struct {
	Svc domain.AdapterServiceHospital
}

func (ce *EchoControllerHospital) CreateHospitalController(c echo.Context) error {
	hospital := model.Hospitals{}
	c.Bind(&hospital)

	err := ce.Svc.CreateHospitalService(hospital)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"Hospitals":    hospital,
	})
}

func (ce *EchoControllerHospital) UpdateHospitalController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	hospital := model.Hospitals{}
	c.Bind(&hospital)

	err := ce.Svc.UpdateHospitalService(intID, hospital)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerHospital) DeleteHospitalController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteHospitalByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerHospital) GetHospitalIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetHospitalsByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"roles":    res,
	})
}

func (ce *EchoControllerHospital) GetHospitalController(c echo.Context) error {
	Hospitalss := ce.Svc.GetAllHospitalsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Hospitalss":    Hospitalss,
	}, "  ")
}