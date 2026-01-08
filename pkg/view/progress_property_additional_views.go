// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ProgressPropertyView ProgressProperty
type ProgressPropertyView struct {
	Progress string `json:"progress,omitempty"`
	Stage    string `json:"stage,omitempty"`
}
