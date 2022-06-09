package handler

import (
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"

	"github.com/labstack/echo/v4"
)

type EchoControllerRole struct {
	svc domain.AdapterServiceRole
}

func (ce *EchoControllerRole) CreateRoleController(c echo.Context) error {
	role := model.Role{}
	c.Bind(&role)

	err := ce.svc.CreateRoleService(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"roles":    role,
	})
}

func (ce *EchoControllerRole) GetRoleIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := ce.svc.GetRoleByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"roles":    res,
	})
}

func (ce *EchoControllerRole) GetAllRoleController(c echo.Context) error {
	roles := ce.svc.GetAllRolesService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"roles":    roles,
	}, "  ")
}

func (ce *EchoControllerRole) DeleteRoleIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteRoleByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
