package theme

import (
	"errors"
	"fmt"
)

const (
	ErrCodeValidation      = "VALIDATION_ERROR"
	ErrCodeNotFound        = "NOT_FOUND"
	ErrCodeCircularRef     = "CIRCULAR_REFERENCE"
	ErrCodeTenantIsolation = "TENANT_ISOLATION"
	ErrCodeStorage         = "STORAGE_ERROR"
	ErrCodeCompilation     = "COMPILATION_ERROR"
	ErrCodeResolution      = "RESOLUTION_ERROR"
	ErrCodeCondition       = "CONDITION_ERROR"
)

type Error struct {
	Code    string
	Message string
	Path    string
	Details map[string]interface{}
	Cause   error
}

func (e *Error) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("[%s] %s (path: %s)", e.Code, e.Message, e.Path)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func (e *Error) WithDetail(key string, value interface{}) *Error {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	e.Details[key] = value
	return e
}

func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Details: make(map[string]interface{}),
	}
}

func NewErrorf(code, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		Details: make(map[string]interface{}),
	}
}

func WrapError(code, message string, cause error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   cause,
		Details: make(map[string]interface{}),
	}
}

var (
	ErrThemeNotFound      = NewError(ErrCodeNotFound, "theme not found")
	ErrInvalidTheme       = NewError(ErrCodeValidation, "invalid theme structure")
	ErrCircularReference  = NewError(ErrCodeCircularRef, "circular token reference detected")
	ErrTenantNotFound     = NewError(ErrCodeNotFound, "tenant not found")
	ErrTenantIsolation    = NewError(ErrCodeTenantIsolation, "tenant isolation violation")
	ErrStorageUnavailable = NewError(ErrCodeStorage, "storage backend unavailable")
	ErrCompilationFailed  = NewError(ErrCodeCompilation, "theme compilation failed")
	ErrResolutionFailed   = NewError(ErrCodeResolution, "token resolution failed")
	ErrConditionFailed    = NewError(ErrCodeCondition, "condition evaluation failed")
)

func IsNotFoundError(err error) bool {
	var themeErr *Error
	if errors.As(err, &themeErr) {
		return themeErr.Code == ErrCodeNotFound
	}
	return false
}

func IsValidationError(err error) bool {
	var themeErr *Error
	if errors.As(err, &themeErr) {
		return themeErr.Code == ErrCodeValidation
	}
	return false
}

func IsCircularReferenceError(err error) bool {
	var themeErr *Error
	if errors.As(err, &themeErr) {
		return themeErr.Code == ErrCodeCircularRef
	}
	return false
}
