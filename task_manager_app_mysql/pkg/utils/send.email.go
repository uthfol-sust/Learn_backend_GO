package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func GenerateEmailCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(1000000))
}

func SendVerificationEmail(clientEmail string) (string, error) {
	userMail := os.Getenv("HOST_EMAIL")
	emailPass := os.Getenv("EMAIL_PASS")
	portStr := os.Getenv("MAIL_PORT")

	code := GenerateEmailCode()

	userPort, err := strconv.Atoi(portStr)
	if err != nil {
		return "", err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", userMail)
	msg.SetHeader("To", clientEmail)
	msg.SetHeader("Subject", "Your Verification Code")
	msg.SetBody("text/plain",
		fmt.Sprintf("Hello!\n\nYour verification code is: %s\n\nIt will expire in 5 minutes.", code))

	d := gomail.NewDialer("smtp.gmail.com", userPort, userMail, emailPass)

	errMsg := d.DialAndSend(msg)

	return code, errMsg
}
