package handlers

import (
	"fmt"
	"go-mailer/models"
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

	to := []string{"viniciusmotta0806@gmail.com"}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "New Message from Contact Form"
	body := `<!DOCTYPE html>
		<html lang="pt">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Contato</title>
				<style>
					#credentials {
						padding: 15px;
						background-color: #71717a;
						border-radius: 5px;
						color: #f4f4f5;
						text-decoration: none;
					}
					
					#credentials p span {
						font-weight: bold;
						color: #d4d4d8;
					}
					
					p {
						color: #f4f4f5;
					}
				</style>
			</head>
			<body style="font-family: Roboto, sans-serif; background-color: #18181b; padding: 20px;">
				<p>Oi Mottinha,</p>
				<p style="padding-bottom: 20px">Alguém quer entrar em contato com você:</p>
				
				<div id="credentials">
					<p><span>Nome:</span> ` + mail.Name + `</p>
					<p><span>Email:</span> ` + mail.Email + `</p>
					<p><span>Mensagem:</span> ` + mail.Message + `</p>
				</div>
				
				<p style="padding-top: 20px">Boa sorte =D</p>
			</body>
		</html>`

	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		from, to, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send email")
	}

	resp := map[string]string{
		"message": "Email sent successfully!",
	}

	return c.JSON(http.StatusOK, resp)
}
