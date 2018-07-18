Mail2
=====

This is a simple tool to send emails from command with an SMTP server. Thus
you could type:

```
$ mail2 -smtp="smtp.orange.fr:25" -from="michel.casabianca@gmail.com" \
  -to="michel.casabianca@gmail.com" -subject="Test" -attach="mail2.go" \
  "This is a test!"
```

This is that simple!

To get help about command line options type :

```
$ mail2 -help
Usage of build/mail2:
  -attach string
      file to attach to the mail
  -bcc string
      coma separated list of blind carbon copy recipients
  -cc string
      coma separated list of carbon copy recipients
  -from string
      mail sender
  -html
      html email message
  -pass string
      user password for authentication
  -smtp string
      smtp server ('smtp.example.com:25')
  -subject string
      mail subject
  -to string
      coma separated list of mail recipients
  -user string
      user name for authentication
```

*Enjoy!*
