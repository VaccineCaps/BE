package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerVStatus struct {
	Svc domain.AdapterServiceVStatus
}

func (ce *EchoControllerVStatus) CreateVStatusController(c echo.Context) error {
	vStatus := model.VaccineStatus{}
	c.Bind(&vStatus)

	err := ce.Svc.CreateVStatusService(vStatus)
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

func (ce *EchoControllerVStatus) GetVStatusIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetVStatusByID(intID)
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

func (ce *EchoControllerVStatus) GetAllVStatusController(c echo.Context) error {
	vStatus := ce.Svc.GetAllVStatusService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"status vaksin":    vStatus,
	}, "  ")
}

func (ce *EchoControllerVStatus) UpdateVStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vStatus := model.VaccineStatus{}
	c.Bind(&vStatus)

	err := ce.Svc.UpdateVStatusService(intID, vStatus)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVStatus) DeleteVStatusIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteVStatusByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
