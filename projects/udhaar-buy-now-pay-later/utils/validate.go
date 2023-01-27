package utils

import (
	"strings"
	"udhaar/constants"

	"github.com/spf13/cast"
)

func IsValidEmail(email string) (string, error) {
	email = strings.ReplaceAll(email, " ", "")
	if len(email) == 0 {
		return "", constants.ErrEmptyEmail
	}
	return email, nil
}

func IsValidName(name string) (string, error) {
	name = strings.ReplaceAll(name, " ", "")
	if len(name) == 0 {
		return "", constants.ErrEmptyName
	}
	return name, nil
}

func IsValidCreditLimit(creditLimitString string) (float64, error) {
	creditLimitString = strings.ReplaceAll(creditLimitString, " ", "")
	if len(creditLimitString) == 0 {
		return 0.0, constants.ErrEmptyCreditLimit
	}

	creditLimitFloat := cast.ToFloat64(creditLimitString)

	if creditLimitFloat <= 0.0 {
		return 0.0, constants.ErrInvalidCreditLimit
	}

	return creditLimitFloat, nil
}
