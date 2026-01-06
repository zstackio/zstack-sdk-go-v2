// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateTemplateConfigParamDetail UpdateTemplateConfig detail param
type UpdateTemplateConfigParamDetail struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateTemplateConfigParam UpdateTemplateConfig request param
type UpdateTemplateConfigParam struct {
	BaseParam
	Params UpdateTemplateConfigParamDetail `json:"params"`
}
// RevertTemplateConfigParamDetail RevertTemplateConfig detail param
type RevertTemplateConfigParamDetail struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// RevertTemplateConfigParam RevertTemplateConfig request param
type RevertTemplateConfigParam struct {
	BaseParam
	Params RevertTemplateConfigParamDetail `json:"params"`
}
// ApplyTemplateConfigParamDetail ApplyTemplateConfig detail param
type ApplyTemplateConfigParamDetail struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// ApplyTemplateConfigParam ApplyTemplateConfig request param
type ApplyTemplateConfigParam struct {
	BaseParam
	Params ApplyTemplateConfigParamDetail `json:"params"`
}
// ResetTemplateConfigParamDetail ResetTemplateConfig detail param
type ResetTemplateConfigParamDetail struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
}

// ResetTemplateConfigParam ResetTemplateConfig request param
type ResetTemplateConfigParam struct {
	BaseParam
	Params ResetTemplateConfigParamDetail `json:"params"`
}
