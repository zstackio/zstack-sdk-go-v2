// Copyright (c) ZStack.io, Inc.

package view

// GetCurrentTimeView GetCurrentTime
type GetCurrentTimeView struct {
	CurrentTime map[string]int64 `json:"currentTime,omitempty"`
	Success bool `json:"success,omitempty"`
}

