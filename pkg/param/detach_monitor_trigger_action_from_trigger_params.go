// Copyright (c) ZStack.io, Inc.

package param

// DetachMonitorTriggerActionFromTriggerDetailParam DetachMonitorTriggerActionFromTrigger detail param
type DetachMonitorTriggerActionFromTriggerDetailParam struct {
	TriggerUuid string `json:"triggerUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
}

// DetachMonitorTriggerActionFromTriggerParam DetachMonitorTriggerActionFromTrigger request param
type DetachMonitorTriggerActionFromTriggerParam struct {
	BaseParam
	Params DetachMonitorTriggerActionFromTriggerDetailParam `json:"params"`
}
