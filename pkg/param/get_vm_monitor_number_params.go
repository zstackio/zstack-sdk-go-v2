// Copyright (c) ZStack.io, Inc.

package param

// GetVmMonitorNumberDetailParam GetVmMonitorNumber detail param
type GetVmMonitorNumberDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmMonitorNumberParam GetVmMonitorNumber request param
type GetVmMonitorNumberParam struct {
	BaseParam
	Params GetVmMonitorNumberDetailParam `json:"params"`
}
