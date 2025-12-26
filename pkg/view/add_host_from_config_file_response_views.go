// Copyright (c) ZStack.io, Inc.

package view

// AddHostFromConfigFileEventView AddHostFromConfigFileEvent
type AddHostFromConfigFileEventView struct {
	Results []AddHostFromFileResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

