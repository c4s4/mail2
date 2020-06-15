# Mail2

- Project :   <https://github.com/c4s4/mail2>.
- Downloads : <https://github.com/c4s4/mail2/releases>.

Mail2 is a tool to send email on command line.

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest Mail2 release with command:

```bash
sh -c "$(curl https://sweetohm.net/dist/mail2/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - https://sweetohm.net/dist/mail2/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install NeON in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/mail2/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *mail2*.

## Usage

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
