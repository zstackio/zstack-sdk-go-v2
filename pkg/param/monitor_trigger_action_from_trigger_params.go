// Copyright (c) ZStack.io, Inc.

package param

// DetachMonitorTriggerActionFromTriggerDetailParam DetachMonitorTriggerActionFromTrigger详细参数
type DetachMonitorTriggerActionFromTriggerDetailParam struct {
	rest string `json:"triggerUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
}

// DetachMonitorTriggerActionFromTriggerParam DetachMonitorTriggerActionFromTrigger请求参数
type DetachMonitorTriggerActionFromTriggerParam struct {
	BaseParam
	Params DetachMonitorTriggerActionFromTriggerDetailParam `json:"params"` // 详细参数
}

