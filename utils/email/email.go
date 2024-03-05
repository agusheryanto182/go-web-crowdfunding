package email

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wneessen/go-mail"
)

type EmailSenderInterface interface {
	QueueEmail(email, otp string) error
}

type Sender struct {
	rdb *redis.Client
}

func NewEmailService(rdb *redis.Client) EmailSenderInterface {
	return &Sender{
		rdb: rdb,
	}
}

func (s *Sender) QueueEmail(email, otp string) error {
	emailData := struct {
		Email string
		OTP   string
	}{
		Email: email,
		OTP:   otp,
	}

	emailBytes, err := json.Marshal(emailData)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := s.rdb.RPush(ctx, "email_queue", emailBytes).Err(); err != nil {
		return err
	}

	return nil
}

// Worker is a function that runs as a backend worker to send emails from a Redis queue
func Worker(rdb *redis.Client) {
	fmt.Println("Starting email worker")
	for {
		emailBytes, err := rdb.LPop(context.Background(), "email_queue").Bytes()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		var emailData struct {
			Email string
			OTP   string
		}
		if err := json.Unmarshal(emailBytes, &emailData); err != nil {
			fmt.Println("Error unmarshalling email data:", err)
			continue
		}

		if err := sendEmail(emailData.Email, emailData.OTP); err != nil {
			fmt.Println("Error sending email:", err)
			continue
		}
		fmt.Println("Email sent successfully")
	}
}

// sendEmail adalah fungsi untuk mengirim email
func sendEmail(email, otp string) error {
	secretUser := os.Getenv("SMTP_USER")
	secretPass := os.Getenv("SMTP_PASS")
	secretPort := os.Getenv("SMTP_PORT")

	convPort, err := strconv.Atoi(secretPort)
	if err != nil {
		return err
	}

	m := mail.NewMsg()
	if err := m.From(secretUser); err != nil {
		return err
	}
	if err := m.To(email); err != nil {
		return err
	}

	m.Subject("Verifikasi Email - Mahasibuk")
	emailTemplate := struct {
		OTP   string
		Email string
	}{
		OTP:   otp,
		Email: email,
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("Failed to get the current file path")
	}

	templatePath := filepath.Join(filepath.Dir(filename), "template.html")

	tmpl, err := template.New("emailTemplate").ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var bodyContent strings.Builder
	if err := tmpl.ExecuteTemplate(&bodyContent, "template.html", emailTemplate); err != nil {
		return err
	}

	m.SetBodyString(mail.TypeTextHTML, bodyContent.String())

	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(convPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(secretUser), mail.WithPassword(secretPass))
	if err != nil {
		return err
	}
	if err := c.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
