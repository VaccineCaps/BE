package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerNews struct {
	Svc domain.AdapterServiceNews
}

func (ce *EchoControllerNews) CreateNewsController(c echo.Context) error {
	News := model.News{}
	c.Bind(&News)

	err := ce.Svc.CreateNewsService(News)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"News":    News,
	})
}

func (ce *EchoControllerNews) UpdateNewsController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	News := model.News{}
	c.Bind(&News)

	err := ce.Svc.UpdateNewsService(intID, News)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerNews) DeleteNewsController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Svc.DeleteNewsByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerNews) GetNewsIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.Svc.GetNewsByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"News":    res,
	})
}


func (ce *EchoControllerNews) GetNewsController(c echo.Context) error {
	News := ce.Svc.GetAllNewsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"News":    News,
	}, "  ")
}