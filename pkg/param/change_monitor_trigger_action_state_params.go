// Copyright (c) ZStack.io, Inc.

package param

// ChangeMonitorTriggerActionStateDetailParam ChangeMonitorTriggerActionState detail param
type ChangeMonitorTriggerActionStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMonitorTriggerActionStateParam ChangeMonitorTriggerActionState request param
type ChangeMonitorTriggerActionStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerActionStateDetailParam `json:"params"`
}
