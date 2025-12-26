// Copyright (c) ZStack.io, Inc.

package param

// CreatePriceTableDetailParam CreatePriceTable detail param
type CreatePriceTableDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Prices []PriceParam `json:"prices" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePriceTableParam CreatePriceTable request param
type CreatePriceTableParam struct {
	BaseParam
	Params CreatePriceTableDetailParam `json:"params"`
}
