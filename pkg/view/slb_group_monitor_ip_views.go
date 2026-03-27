// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SlbGroupMonitorIpInventoryView SlbGroupMonitorIp
type SlbGroupMonitorIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	SlbGroupUuid string `json:"slbGroupUuid,omitempty"`
	MonitorIp string `json:"monitorIp,omitempty"`
}

