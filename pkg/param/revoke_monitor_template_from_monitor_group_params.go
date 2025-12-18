// Copyright (c) ZStack.io, Inc.

package param

// RevokeMonitorTemplateFromMonitorGroupDetailParam RevokeMonitorTemplateFromMonitorGroup详细参数
type RevokeMonitorTemplateFromMonitorGroupDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"templateUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RevokeMonitorTemplateFromMonitorGroupParam RevokeMonitorTemplateFromMonitorGroup请求参数
type RevokeMonitorTemplateFromMonitorGroupParam struct {
	BaseParam
	Params RevokeMonitorTemplateFromMonitorGroupDetailParam `json:"params"` // 详细参数
}

