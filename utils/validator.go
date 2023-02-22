package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/virzz/logger"
)

const (
	ErrMsgArg  = "plz input arg:{%s}"
	ErrMsgFlag = "plz input flag:{%s}"
)

func Validator(field interface{}, tag string, errMsg string) error {
	if err := validator.New().Var(field, tag); err != nil {
		logger.Debug(err)
		return fmt.Errorf(errMsg, tag)
	}
	return nil
}

func ValidArg(field interface{}, tag string) error {
	return Validator(field, tag, ErrMsgArg)
}

func ValidFlag(field interface{}, tag string) error {
	return Validator(field, tag, ErrMsgFlag)
}
