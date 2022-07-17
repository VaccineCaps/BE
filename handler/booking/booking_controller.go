package handler

import (
	"BE/domain"
	"BE/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoControllerBooking struct {
	Svc domain.AdapterServiceBooking
}

func (ce *EchoControllerBooking) CreateBookingHandler(c echo.Context) error {
	booking := model.Booking{}
	c.Bind(&booking)

	err := ce.Svc.CreateBookingService(booking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success add booking",
	})
}

func (ce *EchoControllerBooking) GetAllBookingController(c echo.Context) error {
	Bookings := ce.Svc.GetAllBookingService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"Bookings":    Bookings,
	}, "  ")
}

func (ce *EchoControllerBooking) GetBookingByUserController(c echo.Context) error {
	user_id := c.Param("user_id")
	UserID, err := strconv.Atoi(user_id)
	if err != nil {

	}

	res, err := ce.Svc.GetBookingByUserService(UserID)
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

func (ce *EchoControllerBooking) GetBookingBySessionController(c echo.Context) error {
	session_id := c.Param("session_id")
	SessionID, err := strconv.Atoi(session_id)
	if err != nil {

	}

	res, err := ce.Svc.GetAllBookingBySessionsService(SessionID)
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

func (ce *EchoControllerBooking) GetBookingByIDController(c echo.Context) error {
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
		"bookings":    res,
	})
}

func (ce *EchoControllerBooking) DeleteBookingController(c echo.Context) error {
	book := model.Booking{}
	c.Bind(&book)

	err := ce.Svc.DeleteBookingByIDService(book.UserId, book.HospitalId, book.SessionId, book.CertificateId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
