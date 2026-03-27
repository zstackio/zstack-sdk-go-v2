// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ErrorResultView ErrorResult
type ErrorResultView struct {
	Line int `json:"line,omitempty"`
	Detail string `json:"detail,omitempty"`
}

