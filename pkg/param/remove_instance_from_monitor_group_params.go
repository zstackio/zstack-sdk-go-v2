// Copyright (c) ZStack.io, Inc.

package param

// RemoveInstanceFromMonitorGroupDetailParam RemoveInstanceFromMonitorGroup detail param
type RemoveInstanceFromMonitorGroupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveInstanceFromMonitorGroupParam RemoveInstanceFromMonitorGroup request param
type RemoveInstanceFromMonitorGroupParam struct {
	BaseParam
	Params RemoveInstanceFromMonitorGroupDetailParam `json:"params"`
}
