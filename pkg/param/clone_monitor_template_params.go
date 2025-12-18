// Copyright (c) ZStack.io, Inc.

package param

// CloneMonitorTemplateDetailParam CloneMonitorTemplate detail param
type CloneMonitorTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CloneMonitorTemplateParam CloneMonitorTemplate request param
type CloneMonitorTemplateParam struct {
	BaseParam
	Params CloneMonitorTemplateDetailParam `json:"params"`
}
