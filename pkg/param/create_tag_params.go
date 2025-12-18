// Copyright (c) ZStack.io, Inc.

package param

// CreateTagDetailParam CreateTag detail param
type CreateTagDetailParam struct {
	Name string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
	Description string `json:"description,omitempty"`
	Color string `json:"color,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateTagParam CreateTag request param
type CreateTagParam struct {
	BaseParam
	Params CreateTagDetailParam `json:"params"`
}
