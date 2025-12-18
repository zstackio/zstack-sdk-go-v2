// Copyright (c) ZStack.io, Inc.

package param

// DeleteMonitorTemplateDetailParam DeleteMonitorTemplate detail param
type DeleteMonitorTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTemplateParam DeleteMonitorTemplate request param
type DeleteMonitorTemplateParam struct {
	BaseParam
	Params DeleteMonitorTemplateDetailParam `json:"params"`
}
