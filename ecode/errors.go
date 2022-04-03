package ecode

import (
	"fmt"
	"strconv"
	"strings"
)

type Resp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// Error for error
func (code Resp) Error() string {
	return fmt.Sprintf("%d,%s", code.Code, code.Msg)
}

// Code Code
func (code Resp) Codes() int {
	return code.Code
}

// Message Message
func (code Resp) Message() string {
	return code.Msg
}

// HTTPCode HTTPCode
func (code Resp) HTTPCode() bool {
	return code.Success
}

// Errorf Errorf  v ...interface{}
func Errorf(status bool, code int, msg string) error {
	return &Resp{
		Success: status,
		Data:    nil,
		Code:    code,
		Msg:  msg,
	}
}

// 根据传入的err切割成string
func ErrorString(err error) (int, string) {
	a := strings.FieldsFunc(err.Error(), split)
	b, _ := strconv.Atoi(a[0])
	return b, a[1]
}

func split(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}
