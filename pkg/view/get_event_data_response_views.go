// Copyright (c) ZStack.io, Inc.

package view

// GetEventDataView GetEventData
type GetEventDataView struct {
	Events []EventDataView `json:"events,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

