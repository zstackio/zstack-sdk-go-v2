// Copyright (c) ZStack.io, Inc.

package view

// GetAlarmDataView GetAlarmData
type GetAlarmDataView struct {
	Histories []interface{} `json:"histories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

