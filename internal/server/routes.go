package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.Health)
	e.POST("/send", s.Send)

	return e
}

func (s *Server) Health(c echo.Context) error {
	resp := map[string]string{
		"message": "GO Mailer Server is running",
	}

	return c.JSON(http.StatusOK, resp)
}
