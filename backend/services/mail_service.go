package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/utilities"

)


type EmailService struct {
	config config.SMTPConfig
}


func NewEmailService() *EmailService {
	config := config.GetSMTPConfig()
	return &EmailService{config: config}
}
func (e *EmailService) SendEmail(to []string, subject, body string) error {
	msg := fmt.Sprintf(
		"From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s\r\n",
		e.config.Name, e.config.User, strings.Join(to, ", "), subject, body,
	)

	auth := smtp.PlainAuth("", e.config.User, e.config.Password, e.config.Host)

	serverAddr := fmt.Sprintf("%s:%d", e.config.Host, e.config.Port)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // TODO: Change to false in production
		ServerName:         e.config.Host,
	}

	conn, err := tls.Dial("tcp", serverAddr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, e.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()

	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	if err := client.Mail(e.config.User); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	for _, recipient := range to {
		if err := client.Rcpt(recipient); err != nil {
			utilities.LogError(fmt.Sprintf("Failed to add recipient %s", recipient), err)
			return fmt.Errorf("failed to add recipient %s: %w", recipient, err)
		}
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to start email body write: %w", err)
	}
	if _, err := writer.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write email body: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close email body writer: %w", err)
	}

	// utilities.LogInfo("Email sent successfully to recipients: " + strings.Join(to, ", "))
	return nil
}

func (e *EmailService) PingSMTP() error {
	serverAddr := fmt.Sprintf("%s:%d", e.config.Host, e.config.Port)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,// TODO change to false in production
		ServerName:         e.config.Host,
	}

	conn, err := tls.Dial("tcp", serverAddr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, e.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()

	auth := smtp.PlainAuth("", e.config.User, e.config.Password, e.config.Host)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	return nil
}
