// Copyright (c) ZStack.io, Inc.

package param

// AddPreconfigurationTemplateDetailParam AddPreconfigurationTemplate detail param
type AddPreconfigurationTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Distribution string `json:"distribution" validate:"required"`
	Type string `json:"type" validate:"required"`
	Content string `json:"content" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddPreconfigurationTemplateParam AddPreconfigurationTemplate request param
type AddPreconfigurationTemplateParam struct {
	BaseParam
	Params AddPreconfigurationTemplateDetailParam `json:"params"`
}
