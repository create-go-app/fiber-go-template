package utils

import (
	"bytes"
	"html/template"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

// Sender ...
type Sender struct {
	Login          string
	Password       string
	SMTPServer     string
	SMTPServerPort string
}

// NewEmailSender ...
func NewEmailSender(login, password, server, port string) *Sender {
	return &Sender{login, password, server, port}
}

// SendHTMLEmail ...
func (s *Sender) SendHTMLEmail(template string, dest []string, subject string, data interface{}) error {
	tmpl, errParseTemplate := s.parseTemplate(template, data)
	if errParseTemplate != nil {
		return errParseTemplate
	}
	body := s.writeEmail(dest, "text/html", subject, tmpl)
	s.sendEmail(dest, subject, body) // Send email
	return nil
}

// SendPlainEmail ...
func (s *Sender) SendPlainEmail(dest []string, subject, data string) error {
	body := s.writeEmail(dest, "text/plain", subject, data)
	s.sendEmail(dest, subject, body) // Send email
	return nil
}

// writeEmail ...
func (s *Sender) writeEmail(dest []string, contentType, subject, body string) string {
	//
	header := map[string]string{}
	header["From"] = s.Login
	header["To"] = strings.Join(dest, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = contentType + "; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	//
	var message string
	for key, value := range header {
		message += key + ":" + value + "\r\n"
	}

	//
	var encodedMessage bytes.Buffer
	result := quotedprintable.NewWriter(&encodedMessage)
	result.Write([]byte(body))
	result.Close()

	//
	message += "\r\n" + encodedMessage.String()

	return message
}

// parseTemplate ...
func (s *Sender) parseTemplate(file string, data interface{}) (string, error) {
	//
	tmpl, errParseFiles := template.ParseFiles(file)
	if errParseFiles != nil {
		return "", errParseFiles
	}

	//
	buffer := new(bytes.Buffer)
	if errExecute := tmpl.Execute(buffer, data); errExecute != nil {
		return "", errExecute
	}

	return buffer.String(), nil
}

// sendEmail ...
func (s *Sender) sendEmail(dest []string, subject, body string) error {
	if errSendMail := smtp.SendMail(
		s.SMTPServer+":"+s.SMTPServerPort,
		smtp.PlainAuth("", s.Login, s.Password, s.SMTPServer),
		s.Login, dest, []byte(body),
	); errSendMail != nil {
		return errSendMail
	}
	return nil
}
