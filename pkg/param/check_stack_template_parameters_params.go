// Copyright (c) ZStack.io, Inc.

package param

// CheckStackTemplateParametersDetailParam CheckStackTemplateParameters detail param
type CheckStackTemplateParametersDetailParam struct {
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// CheckStackTemplateParametersParam CheckStackTemplateParameters request param
type CheckStackTemplateParametersParam struct {
	BaseParam
	Params CheckStackTemplateParametersDetailParam `json:"params"`
}
