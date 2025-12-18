// Copyright (c) ZStack.io, Inc.

package param

// UpdateEmailMonitorTriggerActionDetailParam UpdateEmailMonitorTriggerAction detail param
type UpdateEmailMonitorTriggerActionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	MediaUuid string `json:"mediaUuid,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEmailMonitorTriggerActionParam UpdateEmailMonitorTriggerAction request param
type UpdateEmailMonitorTriggerActionParam struct {
	BaseParam
	Params UpdateEmailMonitorTriggerActionDetailParam `json:"params"`
}
