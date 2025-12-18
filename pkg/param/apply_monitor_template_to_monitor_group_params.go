// Copyright (c) ZStack.io, Inc.

package param

// ApplyMonitorTemplateToMonitorGroupDetailParam ApplyMonitorTemplateToMonitorGroup详细参数
type ApplyMonitorTemplateToMonitorGroupDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// ApplyMonitorTemplateToMonitorGroupParam ApplyMonitorTemplateToMonitorGroup请求参数
type ApplyMonitorTemplateToMonitorGroupParam struct {
	BaseParam
	Params ApplyMonitorTemplateToMonitorGroupDetailParam `json:"params"` // 详细参数
}

