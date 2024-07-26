package handlers

import (
	"fmt"
	"go-mailer/models"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"
)

func Send(c echo.Context) error {
	emailAddress := os.Getenv("EMAIL_ADDRESS")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	from := emailAddress
	password := emailPassword

	var mail models.Mail
	if err := c.Bind(&mail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	to := []string{mail.Email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "New Message from Contact Form"
	body := "Name: " + mail.Name + "\nEmail: " + mail.Email + "\nMessage: " + mail.Message

	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		from, to, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send email")
	}

	return echo.NewHTTPError(http.StatusOK, "Email sent successfully!")
}
