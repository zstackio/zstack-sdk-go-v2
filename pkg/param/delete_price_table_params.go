// Copyright (c) ZStack.io, Inc.

package param

// DeletePriceTableDetailParam DeletePriceTable detail param
type DeletePriceTableDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePriceTableParam DeletePriceTable request param
type DeletePriceTableParam struct {
	BaseParam
	Params DeletePriceTableDetailParam `json:"params"`
}
