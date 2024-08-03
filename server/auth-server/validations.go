package main

import (
	"errors"
	"net/mail"
)

func validateUserRegisterData(data RegisterRequestBody) error {
	err := validateEmail(data.Email)

	if err != nil {
		return err
	}

	err = validatePassword(data.Password)

	if err != nil {
		return err
	}

	err = validateUserName(data.Username)

	return err
}

func validateUserName(username string) error {
	if len(username) < 4 {
		return errors.New("username must be at least 4 characters long")
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}