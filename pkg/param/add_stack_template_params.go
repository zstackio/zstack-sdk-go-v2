// Copyright (c) ZStack.io, Inc.

package param

// AddStackTemplateDetailParam AddStackTemplate detail param
type AddStackTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddStackTemplateParam AddStackTemplate request param
type AddStackTemplateParam struct {
	BaseParam
	Params AddStackTemplateDetailParam `json:"params"`
}
