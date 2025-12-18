// Copyright (c) ZStack.io, Inc.

package param

// DeleteResourcePriceDetailParam DeleteResourcePrice detail param
type DeleteResourcePriceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	CutoffPrice bool `json:"cutoffPrice,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteResourcePriceParam DeleteResourcePrice request param
type DeleteResourcePriceParam struct {
	BaseParam
	Params DeleteResourcePriceDetailParam `json:"params"`
}
