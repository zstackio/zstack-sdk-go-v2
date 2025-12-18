// Copyright (c) ZStack.io, Inc.

package param

// UpdatePriceTableDetailParam UpdatePriceTable detail param
type UpdatePriceTableDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePriceTableParam UpdatePriceTable request param
type UpdatePriceTableParam struct {
	BaseParam
	Params UpdatePriceTableDetailParam `json:"params"`
}
