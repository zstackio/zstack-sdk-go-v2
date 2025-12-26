// Copyright (c) ZStack.io, Inc.

package view

// PowerOffHostEventView PowerOffHostEvent
type PowerOffHostEventView struct {
	Results []PowerOffHardwareResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

