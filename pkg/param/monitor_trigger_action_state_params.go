// Copyright (c) ZStack.io, Inc.

package param

// ChangeMonitorTriggerActionStateDetailParam ChangeMonitorTriggerActionState详细参数
type ChangeMonitorTriggerActionStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeMonitorTriggerActionStateParam ChangeMonitorTriggerActionState请求参数
type ChangeMonitorTriggerActionStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerActionStateDetailParam `json:"params"` // 详细参数
}

