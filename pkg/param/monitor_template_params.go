// Copyright (c) ZStack.io, Inc.

package param

// UpdateMonitorTemplateDetailParam UpdateMonitorTemplate详细参数
type UpdateMonitorTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateMonitorTemplateParam UpdateMonitorTemplate请求参数
type UpdateMonitorTemplateParam struct {
	BaseParam
	Params UpdateMonitorTemplateDetailParam `json:"params"` // 详细参数
}

