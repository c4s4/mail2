package main

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail("smtp.orange.fr:25", "michel.casabianca@gmail.com",
		"michel.casabianca@gmail.com", "Test", "This is a test!", "mail2.go")
	if err != nil {
		t.Error("Sending email")
	}
}
