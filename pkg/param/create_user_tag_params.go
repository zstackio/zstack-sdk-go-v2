// Copyright (c) ZStack.io, Inc.

package param

// CreateUserTagDetailParam CreateUserTag detail param
type CreateUserTagDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Tag string `json:"tag" validate:"required"`
}

// CreateUserTagParam CreateUserTag request param
type CreateUserTagParam struct {
	BaseParam
	Params CreateUserTagDetailParam `json:"params"`
}
