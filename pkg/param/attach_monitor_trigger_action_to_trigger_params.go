// Copyright (c) ZStack.io, Inc.

package param

// AttachMonitorTriggerActionToTriggerDetailParam AttachMonitorTriggerActionToTrigger detail param
type AttachMonitorTriggerActionToTriggerDetailParam struct {
	TriggerUuid string `json:"triggerUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
}

// AttachMonitorTriggerActionToTriggerParam AttachMonitorTriggerActionToTrigger request param
type AttachMonitorTriggerActionToTriggerParam struct {
	BaseParam
	Params AttachMonitorTriggerActionToTriggerDetailParam `json:"params"`
}
