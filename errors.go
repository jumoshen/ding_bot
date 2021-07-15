package ding_bot

import (
	"errors"
	"fmt"
)

type ErrCode int

const (
	ErrorCheckUrl ErrCode = iota + 9000
	ErrorDoRequest
	ErrorResponse
)

var e = make(map[ErrCode]string)

func init() {
	e[ErrorCheckUrl] = "check url err:%s"
	e[ErrorDoRequest] = "do request err:%s, body:%s"
	e[ErrorResponse] = "response err:%s, body:%s"
}

func NewError(code ErrCode, params ...interface{}) error {
	errorMessage := DefaultError
	ok := false

	if errorMessage, ok = e[code]; ok && errorMessage != "" {
		errorMessage = fmt.Sprintf(errorMessage, params...)
	}
	return errors.New(errorMessage)
}
