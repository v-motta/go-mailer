package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(c echo.Context) error {
	resp := map[string]string{
		"message": "GO Mailer Server is running",
	}

	return c.JSON(http.StatusOK, resp)
}
