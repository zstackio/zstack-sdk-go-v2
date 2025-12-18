// Copyright (c) ZStack.io, Inc.

package param

// UpdateStackTemplateDetailParam UpdateStackTemplate detail param
type UpdateStackTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State bool `json:"state,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
}

// UpdateStackTemplateParam UpdateStackTemplate request param
type UpdateStackTemplateParam struct {
	BaseParam
	Params UpdateStackTemplateDetailParam `json:"params"`
}
