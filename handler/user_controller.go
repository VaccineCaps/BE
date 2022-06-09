package handler
import (
	"fmt"
	"net/http"
	"strconv"

	"BE/domain"
	"BE/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type EchoControllerUser struct {
	svc domain.AdapterServiceUser
}

func (ce *EchoControllerUser) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	
	err := ce.svc.CreateUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users": user,
	})
}

func (ce *EchoControllerUser) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := ce.svc.UpdateUserService(intID, int(claim["id"].(float64)), user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id": intID,
		"expeted id": claim["id"],
	})
}

func (ce *EchoControllerUser) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerUser) GetUserController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		
	}
	
	res, err := ce.svc.GetUserByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users": res,
	})
}

func (ce *EchoControllerUser) GetUsersController(c echo.Context) error {
	users := ce.svc.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users": users,
	}, "  ")
}

func (ce *EchoControllerUser) LoginUserController(c echo.Context) error {
	userLogin := make(map[string]interface{})

	c.Bind(&userLogin)

	token, statusCode := ce.svc.LoginUser(userLogin["email"].(string), userLogin["password"].(string))
	switch (statusCode) {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "email atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
			return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token": token,
	}, "  ")
}
