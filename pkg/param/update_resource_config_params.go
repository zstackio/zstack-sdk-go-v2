// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourceConfigDetailParam UpdateResourceConfig detail param
type UpdateResourceConfigDetailParam struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateResourceConfigParam UpdateResourceConfig request param
type UpdateResourceConfigParam struct {
	BaseParam
	Params UpdateResourceConfigDetailParam `json:"params"`
}
