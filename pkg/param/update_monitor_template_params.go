// Copyright (c) ZStack.io, Inc.

package param

// UpdateMonitorTemplateDetailParam UpdateMonitorTemplate detail param
type UpdateMonitorTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateMonitorTemplateParam UpdateMonitorTemplate request param
type UpdateMonitorTemplateParam struct {
	BaseParam
	Params UpdateMonitorTemplateDetailParam `json:"params"`
}
