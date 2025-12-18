// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourceStackDetailParam UpdateResourceStack detail param
type UpdateResourceStackDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Rollback bool `json:"rollback,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

// UpdateResourceStackParam UpdateResourceStack request param
type UpdateResourceStackParam struct {
	BaseParam
	Params UpdateResourceStackDetailParam `json:"params"`
}
