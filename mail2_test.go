package main

import (
	"os"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail("smtp.gmail.com:25", "michel.casabianca@gmail.com",
		"michel.casabianca@gmail.com", "", "", "Test", "This is a test!",
		"mail2.go", "michel.casabianca", os.Getenv("MAIL_PASSWORD"), false)
	if err != nil {
		t.Fatalf("Error sending email: %v", err)
	}
}
