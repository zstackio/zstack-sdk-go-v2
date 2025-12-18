// Copyright (c) ZStack.io, Inc.

package param

// DeleteMonitorGroupDetailParam DeleteMonitorGroup detail param
type DeleteMonitorGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMonitorGroupParam DeleteMonitorGroup request param
type DeleteMonitorGroupParam struct {
	BaseParam
	Params DeleteMonitorGroupDetailParam `json:"params"`
}
