package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerCertificate struct {
	Svc domain.AdapterServiceCertificate
}

func (ce *EchoControllerCertificate) CreateCertificateController(c echo.Context) error {
	vStatus := model.Certificate{}
	c.Bind(&vStatus)

	err := ce.Svc.CreateCertificateService(vStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"status vaksin":    vStatus,
	})
}

func (ce *EchoControllerCertificate) GetCertificateIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetCertificateByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"status vaksin":    res,
	})
}

func (ce *EchoControllerCertificate) GetAllCertificateController(c echo.Context) error {
	vStatus := ce.Svc.GetAllCertificateService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"status vaksin":    vStatus,
	}, "  ")
}

func (ce *EchoControllerCertificate) UpdateCertificateController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vStatus := model.Certificate{}
	c.Bind(&vStatus)

	err := ce.Svc.UpdateCertificateService(intID, vStatus)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerCertificate) DeleteCertificateIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteCertificateByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
