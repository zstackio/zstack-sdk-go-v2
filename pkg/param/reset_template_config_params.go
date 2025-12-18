// Copyright (c) ZStack.io, Inc.

package param

// ResetTemplateConfigDetailParam ResetTemplateConfig detail param
type ResetTemplateConfigDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// ResetTemplateConfigParam ResetTemplateConfig request param
type ResetTemplateConfigParam struct {
	BaseParam
	Params ResetTemplateConfigDetailParam `json:"params"`
}
