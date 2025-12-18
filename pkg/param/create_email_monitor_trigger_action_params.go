// Copyright (c) ZStack.io, Inc.

package param

// CreateEmailMonitorTriggerActionDetailParam CreateEmailMonitorTriggerAction detail param
type CreateEmailMonitorTriggerActionDetailParam struct {
	Email string `json:"email" validate:"required"`
	MediaUuid string `json:"mediaUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	TriggerUuids []string `json:"triggerUuids,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEmailMonitorTriggerActionParam CreateEmailMonitorTriggerAction request param
type CreateEmailMonitorTriggerActionParam struct {
	BaseParam
	Params CreateEmailMonitorTriggerActionDetailParam `json:"params"`
}
