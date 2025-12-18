// Copyright (c) ZStack.io, Inc.

package param

// CreateMonitorTemplateDetailParam CreateMonitorTemplate detail param
type CreateMonitorTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorTemplateParam CreateMonitorTemplate request param
type CreateMonitorTemplateParam struct {
	BaseParam
	Params CreateMonitorTemplateDetailParam `json:"params"`
}
