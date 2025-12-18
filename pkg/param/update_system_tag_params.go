// Copyright (c) ZStack.io, Inc.

package param

// UpdateSystemTagDetailParam UpdateSystemTag detail param
type UpdateSystemTagDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Tag string `json:"tag" validate:"required"`
}

// UpdateSystemTagParam UpdateSystemTag request param
type UpdateSystemTagParam struct {
	BaseParam
	Params UpdateSystemTagDetailParam `json:"params"`
}
