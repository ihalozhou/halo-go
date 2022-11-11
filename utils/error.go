package utils

import "errors"

func NewError(err error, message string) error {
	if err != nil {
		println(err.Error())
		// TODO ERR LOG
		return errors.New(message)
	}
	return nil
}
