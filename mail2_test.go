package main

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail("smtp.orange.fr:25", "michel.casabianca@gmail.com",
		"michel.casabianca@gmail.com", "", "", "Test", "This is a test!", "mail2.go", "", "", false, false)
	if err != nil {
		t.Errorf("Sending email: %v", err)
	}
}

func TestSendMailTLS(t *testing.T) {
	err := SendMail("smtp.orange.fr:465", "michel.casabianca@gmail.com",
		"michel.casabianca@gmail.com", "", "", "Test", "This is a test!", "mail2.go",
		"michel_casabianca@orange.fr", "67pwcwh67", false, true)
	if err != nil {
		t.Errorf("Sending email: %v", err)
	}
}
