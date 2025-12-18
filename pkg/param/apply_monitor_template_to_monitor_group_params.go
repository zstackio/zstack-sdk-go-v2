// Copyright (c) ZStack.io, Inc.

package param

// ApplyMonitorTemplateToMonitorGroupDetailParam ApplyMonitorTemplateToMonitorGroup detail param
type ApplyMonitorTemplateToMonitorGroupDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// ApplyMonitorTemplateToMonitorGroupParam ApplyMonitorTemplateToMonitorGroup request param
type ApplyMonitorTemplateToMonitorGroupParam struct {
	BaseParam
	Params ApplyMonitorTemplateToMonitorGroupDetailParam `json:"params"`
}
