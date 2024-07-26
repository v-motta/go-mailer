package routes

import (
	"go-mailer/handlers"

	"github.com/labstack/echo/v4"
)

func Generate(e *echo.Echo) {

	e.GET("/", handlers.Health)
	e.POST("/send", handlers.Send)
}
