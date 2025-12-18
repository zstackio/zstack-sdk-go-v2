// Copyright (c) ZStack.io, Inc.

package param

// CreateResourceStackDetailParam CreateResourceStack detail param
type CreateResourceStackDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Rollback bool `json:"rollback,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourceStackParam CreateResourceStack request param
type CreateResourceStackParam struct {
	BaseParam
	Params CreateResourceStackDetailParam `json:"params"`
}
