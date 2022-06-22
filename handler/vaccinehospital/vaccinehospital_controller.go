package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

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
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	if err != nil {

	}

	res, err := ce.Svc.GetAllStokByHospitalService(HospitalID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"stoks":    res,
	})
}

func (ce *EchoControllerVaccineHospital) GetStokByHospitalVaccineIDController(c echo.Context) error {
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	vaccine_id := c.Param("vaccine_id")
	VaccineID, err := strconv.Atoi(vaccine_id)
	if err != nil {

	}

	res, err := ce.Svc.GetStokByHospitalVaccineService(HospitalID, VaccineID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"stoks":    res,
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
