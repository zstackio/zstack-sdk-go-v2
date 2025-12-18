// Copyright (c) ZStack.io, Inc.

package view

// GetInterfaceServiceTypeStatisticView GetInterfaceServiceTypeStatistic
type GetInterfaceServiceTypeStatisticView struct {
	ServiceTypeStatistics []interface{} `json:"serviceTypeStatistics,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

