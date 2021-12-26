package emailbits

import "github.com/badoux/checkmail"

// Validator válida una dirección de correo electrónico
type Validator interface {
	Validate(email string) error
}

type validator struct {
}

// NewValidator crea un nuevo validador de correos
func NewValidator() Validator {
	return new(validator)
}

func (v *validator) Validate(email string) error {
	return checkmail.ValidateFormat(email)
}
