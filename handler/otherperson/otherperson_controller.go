package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerOther struct {
	Svc domain.AdapterServiceOther
}

func (ce *EchoControllerOther) CreateOtherController(c echo.Context) error {
	Other := model.OtherPerson{}
	c.Bind(&Other)

	err := ce.Svc.CreateOtherService(Other)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"OtherPerson":    Other,
	})
}

func (ce *EchoControllerOther) UpdateOtherController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	Other := model.OtherPerson{}
	c.Bind(&Other)

	err := ce.Svc.UpdateOtherService(intID, Other)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerOther) DeleteOtherController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteOtherByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerOther) GetOtherIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetOtherByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"OtherPerson":    res,
	})
}


func (ce *EchoControllerOther) GetOtherController(c echo.Context) error {
	OtherPersons := ce.Svc.GetAllOtherService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"OtherPersons":    OtherPersons,
	}, "  ")
}