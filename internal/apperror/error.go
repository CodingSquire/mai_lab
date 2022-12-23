package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrorNotFound = NewAppError("not found", "", "code-0")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal

}

func NewAppError(message, developerMessage, code string) *AppError {

	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, "NS-000002", "some thing wrong with user data")
}

func systemError(err error) *AppError {
	return NewAppError("system error", err.Error(), "US-000000")
}
