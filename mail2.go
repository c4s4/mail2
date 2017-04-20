//usr/bin/env go run $0 "$@" ; exit

package main

import (
	"flag"
	"fmt"
	"github.com/jordan-wright/email"
	"strings"
)

func SendMail(smtp, from, to, subject, text, attach string) error {
	mail := email.NewEmail()
	mail.From = from
	mail.To = strings.Split(to, ",")
	mail.Subject = subject
	mail.Text = []byte(text)
	if attach != "" {
		mail.AttachFile(attach)
	}
	return mail.Send(smtp, nil)
}

func main() {
	smtp := flag.String("smtp", "", "smtp server ('smtp.example.com:25')")
	from := flag.String("from", "", "mail sender")
	to := flag.String("to", "", "mail recipient")
	subject := flag.String("subject", "", "mail subject")
	attach := flag.String("attach", "", "file to attach to the mail")
	flag.Parse()
	text := strings.Join(flag.Args(), " ")
	err := SendMail(*smtp, *from, *to, *subject, text, *attach)
	if err != nil {
		println(fmt.Sprintf("ERROR sending mail: %v", err))
	}
}
