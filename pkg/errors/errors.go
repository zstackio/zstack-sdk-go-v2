// Copyright (c) ZStack.io, Inc.

package errors

import "fmt"

// Error custom error type
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotFound    = Error("NotFoundError")
	ErrDuplicateId = Error("DuplicateIdError")
	ErrParameter   = Error("ParameterError")
	ErrAuth        = Error("AuthError")
	ErrPermission  = Error("PermissionError")
	ErrInternal    = Error("InternalError")
)

// Wrap wraps an error with a message
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// ErrorCode API error code structure
type ErrorCode struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	Details     string     `json:"details"`
	Elaboration string     `json:"elaboration"`
	Cause       *ErrorCode `json:"cause"`
}
