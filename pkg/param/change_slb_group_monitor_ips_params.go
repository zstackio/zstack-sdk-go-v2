// Copyright (c) ZStack.io, Inc.

package param

// ChangeSlbGroupMonitorIpsDetailParam ChangeSlbGroupMonitorIps detail param
type ChangeSlbGroupMonitorIpsDetailParam struct {
	SlbGroupUuid string `json:"slbGroupUuid" validate:"required"`
	MonitorIps []string `json:"monitorIps" validate:"required"`
}

// ChangeSlbGroupMonitorIpsParam ChangeSlbGroupMonitorIps request param
type ChangeSlbGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeSlbGroupMonitorIpsDetailParam `json:"params"`
}
