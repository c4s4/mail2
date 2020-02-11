package main

import (
	"os"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail(os.Getenv("MAIL_SMTP"),
					os.Getenv("MAIL_RECIPIENT"),
					os.Getenv("MAIL_RECIPIENT"),
					"", "", "Test", "This is a test!", "mail2.go",
					os.Getenv("MAIL_USERNAME"),
					os.Getenv("MAIL_PASSWORD"),
					false)
	if err != nil {
		t.Fatalf("Error sending email: %v", err)
	}
}
