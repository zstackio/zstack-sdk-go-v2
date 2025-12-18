// Copyright (c) ZStack.io, Inc.

package param

// UpdateTemplateConfigDetailParam UpdateTemplateConfig detail param
type UpdateTemplateConfigDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateTemplateConfigParam UpdateTemplateConfig request param
type UpdateTemplateConfigParam struct {
	BaseParam
	Params UpdateTemplateConfigDetailParam `json:"params"`
}
