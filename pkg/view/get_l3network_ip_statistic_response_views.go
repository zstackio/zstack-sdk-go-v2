// Copyright (c) ZStack.io, Inc.

package view

// GetL3NetworkIpStatisticView GetL3NetworkIpStatistic
type GetL3NetworkIpStatisticView struct {
	IpStatistics []IpStatisticDataView `json:"ipStatistics,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

