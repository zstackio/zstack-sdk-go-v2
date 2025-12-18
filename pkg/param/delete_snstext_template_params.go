// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSTextTemplateDetailParam DeleteSNSTextTemplate detail param
type DeleteSNSTextTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSTextTemplateParam DeleteSNSTextTemplate request param
type DeleteSNSTextTemplateParam struct {
	BaseParam
	Params DeleteSNSTextTemplateDetailParam `json:"params"`
}
