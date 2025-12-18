// Copyright (c) ZStack.io, Inc.

package param

// UpdateEmailMonitorTriggerActionDetailParam UpdateEmailMonitorTriggerAction详细参数
type UpdateEmailMonitorTriggerActionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"email,omitempty"`
	rest string `json:"mediaUuid,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateEmailMonitorTriggerActionParam UpdateEmailMonitorTriggerAction请求参数
type UpdateEmailMonitorTriggerActionParam struct {
	BaseParam
	Params UpdateEmailMonitorTriggerActionDetailParam `json:"params"` // 详细参数
}

