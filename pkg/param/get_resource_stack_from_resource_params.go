// Copyright (c) ZStack.io, Inc.

package param

// GetResourceStackFromResourceDetailParam GetResourceStackFromResource detail param
type GetResourceStackFromResourceDetailParam struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// GetResourceStackFromResourceParam GetResourceStackFromResource request param
type GetResourceStackFromResourceParam struct {
	BaseParam
	Params GetResourceStackFromResourceDetailParam `json:"params"`
}
