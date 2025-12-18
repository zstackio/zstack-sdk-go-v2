// Copyright (c) ZStack.io, Inc.

package param

// ChangeMonitorTriggerStateDetailParam ChangeMonitorTriggerState详细参数
type ChangeMonitorTriggerStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeMonitorTriggerStateParam ChangeMonitorTriggerState请求参数
type ChangeMonitorTriggerStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerStateDetailParam `json:"params"` // 详细参数
}

