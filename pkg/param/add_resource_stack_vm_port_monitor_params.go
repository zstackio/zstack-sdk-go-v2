// Copyright (c) ZStack.io, Inc.

package param

// AddResourceStackVmPortMonitorDetailParam AddResourceStackVmPortMonitor detail param
type AddResourceStackVmPortMonitorDetailParam struct {
	StackUuid string `json:"stackUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Port int `json:"port" validate:"required"`
}

// AddResourceStackVmPortMonitorParam AddResourceStackVmPortMonitor request param
type AddResourceStackVmPortMonitorParam struct {
	BaseParam
	Params AddResourceStackVmPortMonitorDetailParam `json:"params"`
}
