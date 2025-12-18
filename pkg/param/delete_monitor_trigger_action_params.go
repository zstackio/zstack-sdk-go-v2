// Copyright (c) ZStack.io, Inc.

package param

// DeleteMonitorTriggerActionDetailParam DeleteMonitorTriggerAction detail param
type DeleteMonitorTriggerActionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTriggerActionParam DeleteMonitorTriggerAction request param
type DeleteMonitorTriggerActionParam struct {
	BaseParam
	Params DeleteMonitorTriggerActionDetailParam `json:"params"`
}
