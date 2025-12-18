// Copyright (c) ZStack.io, Inc.

package param

// ApplyTemplateConfigDetailParam ApplyTemplateConfig detail param
type ApplyTemplateConfigDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// ApplyTemplateConfigParam ApplyTemplateConfig request param
type ApplyTemplateConfigParam struct {
	BaseParam
	Params ApplyTemplateConfigDetailParam `json:"params"`
}
