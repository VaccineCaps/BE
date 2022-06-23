package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerSession struct {
	Svc domain.AdapterServiceSession
}

func (ce *EchoControllerSession) CreateSessionHandler(c echo.Context) error {
	session := model.Session{}
	c.Bind(&session)

	err := ce.Svc.CreateSessionService(session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success add session",
	})
}

func (ce *EchoControllerSession) UpdateVaccineSessionController(c echo.Context) error {
	session := model.Session{}
	c.Bind(&session)

	err := ce.Svc.UpdateSessionService(session.HospitalId, session.VaccineId, session)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerSession) GetSessionByHospitalController(c echo.Context) error {
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	if err != nil {

	}

	res, err := ce.Svc.GetAllSessionByHospitalService(HospitalID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no session available",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Sessions": res,
	})
}

func (ce *EchoControllerSession) GetSessionByHospitalVaccineIDController(c echo.Context) error {
	hospital_id := c.Param("hospital_id")
	HospitalID, err := strconv.Atoi(hospital_id)
	vaccine_id := c.Param("vaccine_id")
	VaccineID, err := strconv.Atoi(vaccine_id)
	if err != nil {

	}

	res, err := ce.Svc.GetSessionByHospitalVaccineService(HospitalID, VaccineID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no session available",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"sessions": res,
	})
}

func (ce *EchoControllerSession) DeleteSessionIDController(c echo.Context) error {
	session := model.Session{}
	c.Bind(&session)

	err := ce.Svc.DeleteSessionByIDService(session.HospitalId, session.VaccineId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
