// Copyright (c) ZStack.io, Inc.

package param

// RevertTemplateConfigDetailParam RevertTemplateConfig detail param
type RevertTemplateConfigDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// RevertTemplateConfigParam RevertTemplateConfig request param
type RevertTemplateConfigParam struct {
	BaseParam
	Params RevertTemplateConfigDetailParam `json:"params"`
}
