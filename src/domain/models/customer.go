package domain

import (
	"errors"
	"net/mail"

	"github.com/mvrilo/go-cpf"
)

type Customer struct {
	Name           string
	Email          string
	DocumentNumber string
}

func validateDocumentNumber(documentNumber string) error {
	_, err := cpf.Valid(documentNumber)
	if documentNumber == "" || err != nil {
		return errors.New("invalid document number")
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if email == "" || err != nil {
		return errors.New("invalid email")
	}

	return nil
}

func validateName(email string) error {
	if email == "" {
		return errors.New("invalid name")
	}

	return nil
}

func (customer Customer) Validate() error {

	if err := validateDocumentNumber(customer.DocumentNumber); err != nil {
		return err
	}

	if err := validateEmail(customer.Email); err != nil {
		return err
	}

	if err := validateName(customer.Name); err != nil {
		return err
	}

	return nil
}
