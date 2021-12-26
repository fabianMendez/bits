package emailbits

import "fmt"

// Config holds the parameters needed to send emails
type Config struct {
	Host     string `envconfig:"email_host"`
	Port     int    `envconfig:"email_port"`
	Username string `envconfig:"email_username"`
	Password string `envconfig:"email_password"`
}

// Validate checks if the email configuration is valid
func (e *Config) Validate() error {
	if e.Host == "" {
		return fmt.Errorf("host is required")
	}

	if e.Port == 0 {
		return fmt.Errorf("port is required")
	}

	if e.Username == "" {
		return fmt.Errorf("username is required")
	}

	if e.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}
