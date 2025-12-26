// Copyright (c) ZStack.io, Inc.

package view

// GetActiveAlarmStatusView GetActiveAlarmStatus
type GetActiveAlarmStatusView struct {
	Statuses []ActiveAlarmStatusView `json:"statuses,omitempty"`
	Success bool `json:"success,omitempty"`
}

