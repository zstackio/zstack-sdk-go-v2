// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ErrorCodeView ErrorCode
type ErrorCodeView struct {
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Details string `json:"details,omitempty"`
	Elaboration string `json:"elaboration,omitempty"`
	Location string `json:"location,omitempty"`
	Cost string `json:"cost,omitempty"`
	Cause *ErrorCodeView `json:"cause,omitempty"`
	Opaque interface{} `json:"opaque,omitempty"`
	GlobalErrorCode string `json:"globalErrorCode,omitempty"`
	Message string `json:"message,omitempty"`
}

