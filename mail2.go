//usr/bin/env go run $0 "$@" ; exit

package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/jordan-wright/email"
)

// List makes a slice from a coma separated list:
// - list: coma separated list
// Return: list as a slice.
func List(list string) []string {
	result := strings.Split(list, ",")
	for index, element := range result {
		result[index] = strings.TrimSpace(element)
	}
	return result
}

// SendMail sends an email:
// - server: the SMTP server address (such as 'smtp.orange.fr:25')
// - from: the email sender (such as 'test@example.com').
// - to: the email recipient(s) as a coma separated list.
// - cc: carbon copy recipient(s) as a coma separated list.
// - bcc: blind carbon copy recipient(s) as a coma separated list.
// - subject: the email subject.
// - text: the email text.
// - attach: email attachment as a file path.
// - user: SMTP server user name.
// - password: the SMTP server password.
// - html: tells if email is HTML.
func SendMail(server, from, to, cc, bcc, subject, text, attach, user, pass string, html bool) error {
	mail := email.NewEmail()
	mail.From = from
	mail.To = List(to)
	if cc != "" {
		mail.Cc = List(cc)
	}
	if bcc != "" {
		mail.Bcc = List(bcc)
	}
	mail.Subject = subject
	if html {
		mail.HTML = []byte(text)
	} else {
		mail.Text = []byte(text)
	}
	if attach != "" {
		if _, err := os.Stat(attach); os.IsNotExist(err) {
			return fmt.Errorf("attachement file '%s' not found", attach)
		}
		mail.AttachFile(attach)
	}
	var auth smtp.Auth
	if user != "" {
		index := strings.LastIndex(server, ":")
		if index == -1 {
			return fmt.Errorf("smtp must be like hostname:port, no port found")
		}
		host := server[:index]
		auth = smtp.PlainAuth("", user, pass, host)
	}
	return mail.Send(server, auth)
}

func main() {
	smtp := flag.String("smtp", "", "smtp server ('smtp.example.com:25')")
	from := flag.String("from", "", "mail sender")
	to := flag.String("to", "", "coma separated list of mail recipients")
	cc := flag.String("cc", "", "coma separated list of carbon copy recipients")
	bcc := flag.String("bcc", "", "coma separated list of blind carbon copy recipients")
	subject := flag.String("subject", "", "mail subject")
	html := flag.Bool("html", false, "html email message")
	attach := flag.String("attach", "", "file to attach to the mail")
	user := flag.String("user", "", "user name for authentication")
	pass := flag.String("pass", "", "user password for authentication")
	flag.Parse()
	text := strings.Join(flag.Args(), " ")
	err := SendMail(*smtp, *from, *to, *cc, *bcc, *subject, text, *attach, *user, *pass, *html)
	if err != nil {
		println(fmt.Sprintf("ERROR sending mail: %v", err))
		os.Exit(1)
	}
}
