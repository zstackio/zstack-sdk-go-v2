// Copyright (c) ZStack.io, Inc.

package param

// UpdatePreconfigurationTemplateDetailParam UpdatePreconfigurationTemplate detail param
type UpdatePreconfigurationTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

// UpdatePreconfigurationTemplateParam UpdatePreconfigurationTemplate request param
type UpdatePreconfigurationTemplateParam struct {
	BaseParam
	Params UpdatePreconfigurationTemplateDetailParam `json:"params"`
}
