// Copyright (c) ZStack.io, Inc.

package param

// ChangeVpcHaGroupMonitorIpsDetailParam ChangeVpcHaGroupMonitorIps detail param
type ChangeVpcHaGroupMonitorIpsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonitorIps []string `json:"monitorIps,omitempty"`
}

// ChangeVpcHaGroupMonitorIpsParam ChangeVpcHaGroupMonitorIps request param
type ChangeVpcHaGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeVpcHaGroupMonitorIpsDetailParam `json:"params"`
}
