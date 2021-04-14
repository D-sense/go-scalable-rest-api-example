package errors

import (
	"encoding/json"
)

type CustomError struct {
	code string
	err  string
}

func (e *CustomError) Code() string {
	if e == nil {
		return ""
	}
	return e.code
}

func (e *CustomError) Error() string {
	if e == nil {
		return ""
	}
	return e.err
}

func (e CustomError) MarshalJsonResponse() ([]byte, error) {
	j := struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}{
		Code:    e.code,
		Message: e.Error(),
	}
	return json.Marshal(&j)
}
