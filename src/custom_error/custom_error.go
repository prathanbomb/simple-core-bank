package custom_error

import (
	"strings"
)

type ValidationError struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

type AuthorizationError struct {
	Code           uint64 `json:"code"`
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

type UserError struct {
	Code           uint64 `json:"code"`
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
}

func (e *UserError) Error() string {
	return e.Message
}

type InternalError struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

func (e *InternalError) Error() string {
	return e.Message
}

type ListErr []error

func (e ListErr) Error() string {
	var l []string
	for _, v := range e {
		if v == nil {
			continue
		}

		l = append(l, v.Error())
	}
	return strings.Join(l, " ,")
}
