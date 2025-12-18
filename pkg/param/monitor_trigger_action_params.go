// Copyright (c) ZStack.io, Inc.

package param

// DeleteMonitorTriggerActionDetailParam DeleteMonitorTriggerAction详细参数
type DeleteMonitorTriggerActionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTriggerActionParam DeleteMonitorTriggerAction请求参数
type DeleteMonitorTriggerActionParam struct {
	BaseParam
	Params DeleteMonitorTriggerActionDetailParam `json:"params"` // 详细参数
}

