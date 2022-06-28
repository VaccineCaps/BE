package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerAdvertise struct {
	Svc domain.AdapterServiceAdvertise
}

func (ce *EchoControllerAdvertise) CreateAdvertiseController(c echo.Context) error {
	Advertise := model.Advertise{}
	c.Bind(&Advertise)

	err := ce.Svc.CreateAdvertiseService(Advertise)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"Advertise":    Advertise,
	})
}

func (ce *EchoControllerAdvertise) UpdateAdvertiseController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	Advertise := model.Advertise{}
	c.Bind(&Advertise)

	err := ce.Svc.UpdateAdvertiseService(intID, Advertise)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerAdvertise) DeleteAdvertiseController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteAdvertiseByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerAdvertise) GetAdvertiseIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetAdvertiseByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Advertise":    res,
	})
}


func (ce *EchoControllerAdvertise) GetAdvertiseController(c echo.Context) error {
	Advertise := ce.Svc.GetAllAdvertiseService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Advertise":    Advertise,
	}, "  ")
}