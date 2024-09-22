package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password),
	)
}

func ValidateStruct(s any) error {
	validate := validator.New()
	err := validate.Struct(s)
	errMsgs := make([]string, 0)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"error when validating %s: '%v'",
				err.Field(),
				err.Value(),
			))
		}
		return errors.New(errMsgs[0])
	}
	return nil
}
