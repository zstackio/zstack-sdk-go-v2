// Copyright (c) ZStack.io, Inc.

package view

// GetActiveAlarmStatusView GetActiveAlarmStatus
type GetActiveAlarmStatusView struct {
	Statuses []interface{} `json:"statuses,omitempty"`
	Success bool `json:"success,omitempty"`
}

