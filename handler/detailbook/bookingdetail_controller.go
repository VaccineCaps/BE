package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerBookingDetail struct {
	Svc domain.AdapterServiceBookingDetail
}

func (ce *EchoControllerBookingDetail) CreateBookingDetailHandler(c echo.Context) error {
	booking := model.BookingDetail{}
	c.Bind(&booking)

	err := ce.Svc.CreateBookingDetailService(booking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success add booking detail",
	})
}

func (ce *EchoControllerBookingDetail) GetAllBookingDetailController(c echo.Context) error {
	BookingDetails := ce.Svc.GetAllDetailService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"BookingDetails":    BookingDetails,
	}, "  ")
}

func (ce *EchoControllerBookingDetail) GetBookingDetailByUserController(c echo.Context) error {
	user_id := c.Param("user_id")
	UserID, err := strconv.Atoi(user_id)
	if err != nil {

	}

	res, err := ce.Svc.GetBookingDetailByUserService(UserID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"bookings":    res,
	})
}

func (ce *EchoControllerBookingDetail) GetBookingDetailByOPController(c echo.Context) error {
	otherperson_id := c.Param("otherperson_id")
	OPID, err := strconv.Atoi(otherperson_id)
	if err != nil {

	}

	res, err := ce.Svc.GetBookingDetailByOtherPersonService(OPID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"booking detail":    res,
	})
}

func (ce *EchoControllerBookingDetail) GetBookingDetailByIDController(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.Svc.GetBookingsByIDService(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"bookings details":    res,
	})
}

func (ce *EchoControllerBookingDetail) GetBookingDetailByBookingController(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.Svc.GetBookingDetailByBookingsService(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"bookings details":    res,
	})
}

