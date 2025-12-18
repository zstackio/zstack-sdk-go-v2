// Copyright (c) ZStack.io, Inc.

package param

// CreateSystemTagDetailParam CreateSystemTag detail param
type CreateSystemTagDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Tag string `json:"tag" validate:"required"`
}

// CreateSystemTagParam CreateSystemTag request param
type CreateSystemTagParam struct {
	BaseParam
	Params CreateSystemTagDetailParam `json:"params"`
}
