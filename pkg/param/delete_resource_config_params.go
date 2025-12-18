// Copyright (c) ZStack.io, Inc.

package param

// DeleteResourceConfigDetailParam DeleteResourceConfig detail param
type DeleteResourceConfigDetailParam struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteResourceConfigParam DeleteResourceConfig request param
type DeleteResourceConfigParam struct {
	BaseParam
	Params DeleteResourceConfigDetailParam `json:"params"`
}
