// The appErrors package provides support for application specific error handling
package appErrors

import (
	"fmt"
)

const (
	VALIDATION_ERROR_MSG  = "Validation Error: The parameters provided were invalid."
	VALIDATION_ERROR_CODE = 409

	UNAUTHORIZED_ERROR_MSG  = "Invalid Credentials were supplied."
	UNAUTHORIZED_ERROR_CODE = 401

	APP_ERROR_MSG  = "An Application Error has occured."
	APP_ERROR_CODE = 500

	CACHE_CONTROL_HEADER = "Cache-control"
)

type (
	// AppError implements an application error which
	// reguires an error string and code
	AppError struct {
		ErrorMsg string
		Code     int
	}
)

// Error returns the error message that is associated with the AppError object
func (this *AppError) Error() string {
	return this.ErrorMsg
}

// ErrorCode returns the error code that is associated with the AppError object
func (this *AppError) ErrorCode() int {
	if this.Code == 0 {
		return APP_ERROR_CODE
	}

	return this.Code
}

// NewError creates an AppError object
func NewError(error string, code int) error {
	return &AppError{ErrorMsg: error, Code: code}
}

// NewValidationError creates an AppError object for a validation error
func NewValidationError(error string) error {
	return &AppError{ErrorMsg: error, Code: VALIDATION_ERROR_CODE}
}

// ErrorValf creates an AppError based validation error from the formatted string
func ErrorValf(format string, a ...interface{}) error {
	return NewError(fmt.Sprintf(format, a), VALIDATION_ERROR_CODE)
}