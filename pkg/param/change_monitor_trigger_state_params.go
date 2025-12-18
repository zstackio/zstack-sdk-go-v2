// Copyright (c) ZStack.io, Inc.

package param

// ChangeMonitorTriggerStateDetailParam ChangeMonitorTriggerState detail param
type ChangeMonitorTriggerStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMonitorTriggerStateParam ChangeMonitorTriggerState request param
type ChangeMonitorTriggerStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerStateDetailParam `json:"params"`
}
