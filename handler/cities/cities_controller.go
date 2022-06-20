package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerCity struct {
	Svc domain.AdapterServiceCity
}

func (ce *EchoControllerCity) CreateCityController(c echo.Context) error {
	City := model.Cities{}
	c.Bind(&City)

	err := ce.Svc.CreateCityService(City)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"Cities":    City,
	})
}

func (ce *EchoControllerCity) GetCityIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetCityByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Cities":    res,
	})
}

func (ce *EchoControllerCity) GetAllCityController(c echo.Context) error {
	Cities := ce.Svc.GetAllCityService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Cities":    Cities,
	}, "  ")
}

func (ce *EchoControllerCity) DeleteCityIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteCityByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
