// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcHaGroupMonitorIpInventoryView VpcHaGroupMonitorIp
type VpcHaGroupMonitorIpInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vpcHaRouterUuid,omitempty"`
	rest string `json:"monitorIp,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

