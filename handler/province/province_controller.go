package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerProvince struct {
	Svc domain.AdapterServiceProvince
}

func (ce *EchoControllerProvince) CreateProvinceController(c echo.Context) error {
	province := model.Provinces{}
	c.Bind(&province)

	err := ce.Svc.CreateProvinceService(province)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"provinces":    province,
	})
}

func (ce *EchoControllerProvince) GetProvinceIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetProvinceByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"provinces":    res,
	})
}

func (ce *EchoControllerProvince) GetAllProvinceController(c echo.Context) error {
	provinces := ce.Svc.GetAllProvinceService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"provinces":    provinces,
	}, "  ")
}

func (ce *EchoControllerProvince) DeleteProvinceIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteProvinceByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
