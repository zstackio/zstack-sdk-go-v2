// Copyright (c) ZStack.io, Inc.

package param

// DeleteResourceStackVmPortMonitorDetailParam DeleteResourceStackVmPortMonitor detail param
type DeleteResourceStackVmPortMonitorDetailParam struct {
	StackUuid string `json:"stackUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Port int `json:"port,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteResourceStackVmPortMonitorParam DeleteResourceStackVmPortMonitor request param
type DeleteResourceStackVmPortMonitorParam struct {
	BaseParam
	Params DeleteResourceStackVmPortMonitorDetailParam `json:"params"`
}
