// Copyright (c) ZStack.io, Inc.

package param

// AttachMonitorTriggerActionToTriggerDetailParam AttachMonitorTriggerActionToTrigger详细参数
type AttachMonitorTriggerActionToTriggerDetailParam struct {
	rest string `json:"triggerUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
}

// AttachMonitorTriggerActionToTriggerParam AttachMonitorTriggerActionToTrigger请求参数
type AttachMonitorTriggerActionToTriggerParam struct {
	BaseParam
	Params AttachMonitorTriggerActionToTriggerDetailParam `json:"params"` // 详细参数
}

