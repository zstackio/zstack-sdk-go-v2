// Copyright (c) ZStack.io, Inc.

package param

// GetResourceConfigDetailParam GetResourceConfig detail param
type GetResourceConfigDetailParam struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// GetResourceConfigParam GetResourceConfig request param
type GetResourceConfigParam struct {
	BaseParam
	Params GetResourceConfigDetailParam `json:"params"`
}
