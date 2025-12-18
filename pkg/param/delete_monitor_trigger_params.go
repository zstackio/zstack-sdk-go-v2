// Copyright (c) ZStack.io, Inc.

package param

// DeleteMonitorTriggerDetailParam DeleteMonitorTrigger detail param
type DeleteMonitorTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTriggerParam DeleteMonitorTrigger request param
type DeleteMonitorTriggerParam struct {
	BaseParam
	Params DeleteMonitorTriggerDetailParam `json:"params"`
}
