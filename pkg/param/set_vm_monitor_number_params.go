// Copyright (c) ZStack.io, Inc.

package param

// SetVmMonitorNumberDetailParam SetVmMonitorNumber detail param
type SetVmMonitorNumberDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonitorNumber int `json:"monitorNumber" validate:"required"`
}

// SetVmMonitorNumberParam SetVmMonitorNumber request param
type SetVmMonitorNumberParam struct {
	BaseParam
	Params SetVmMonitorNumberDetailParam `json:"params"`
}
