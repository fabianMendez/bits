package emailbits

import (
	"bytes"
	"text/template"

	"gopkg.in/gomail.v2"
)

// Service permite enviar correos electrónicos
type Service interface {
	Send(subject, body string, data interface{}, to ...string) error
}

type service struct {
	from   string
	dialer *gomail.Dialer
}

// NewService crea un nuevo servicio de envío de correos
func NewService(c *Config) (Service, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	return &service{
		from:   c.Username,
		dialer: gomail.NewDialer(c.Host, c.Port, c.Username, c.Password),
	}, nil
}

func (s *service) Send(subject, body string, data interface{}, to ...string) error {
	buf := bytes.NewBuffer(nil)
	t := template.Must(template.New("email").Parse(body))

	err := t.Execute(buf, data)
	if err != nil {
		return err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", s.from)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", buf.String())

	return s.dialer.DialAndSend(message)
}
