// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SlbGroupMonitorIpInventoryView SlbGroupMonitorIp
type SlbGroupMonitorIpInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"slbGroupUuid,omitempty"`
	rest string `json:"monitorIp,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

