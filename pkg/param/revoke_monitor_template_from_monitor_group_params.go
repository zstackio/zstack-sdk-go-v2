// Copyright (c) ZStack.io, Inc.

package param

// RevokeMonitorTemplateFromMonitorGroupDetailParam RevokeMonitorTemplateFromMonitorGroup detail param
type RevokeMonitorTemplateFromMonitorGroupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	TemplateUuid string `json:"templateUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RevokeMonitorTemplateFromMonitorGroupParam RevokeMonitorTemplateFromMonitorGroup request param
type RevokeMonitorTemplateFromMonitorGroupParam struct {
	BaseParam
	Params RevokeMonitorTemplateFromMonitorGroupDetailParam `json:"params"`
}
