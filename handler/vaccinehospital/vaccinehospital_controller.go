package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoControllerVaccineHospital struct {
	Svc domain.AdapterServiceVaccineHospital
}

func (ce *EchoControllerVaccineHospital) CreateStokHandler(c echo.Context) error {
	stok := model.VaccineHospitals{}
	c.Bind(&stok)

	err := ce.Svc.CreateStokService(stok)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success add stok",
	})
}

func (ce *EchoControllerVaccineHospital) UpdateVaccineStokController(c echo.Context) error {
	stok := model.VaccineHospitals{}
	c.Bind(&stok)

	err := ce.Svc.UpdateVaccineStokService(stok.HospitalId, stok.VaccineId, stok)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccineHospital) GetStokByHospitalController(c echo.Context) error {
	stok := model.VaccineHospitals{}
	c.Bind(&stok)

	res, err := ce.Svc.GetAllStokByHospitalService(stok.HospitalId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

func (ce *EchoControllerVaccineHospital) GetStokByHospitalVaccineIDController(c echo.Context) error {
	stok := model.VaccineHospitals{}
	c.Bind(&stok)

	res, err := ce.Svc.GetStokByHospitalVaccineService(stok.HospitalId, stok.VaccineId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

func (ce *EchoControllerVaccineHospital) DeleteHospitalVaccineIDController(c echo.Context) error {
	stok := model.VaccineHospitals{}
	c.Bind(&stok)

	err := ce.Svc.DeleteVaccineStokByIDService(stok.HospitalId, stok.VaccineId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
