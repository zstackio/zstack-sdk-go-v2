// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreatePriceTableParamDetail CreatePriceTable detail param
type CreatePriceTableParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Prices []CreatePriceTable_PriceParam `json:"prices" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePriceTableParam CreatePriceTable request param
type CreatePriceTableParam struct {
	BaseParam
	CreatePriceTable CreatePriceTableParamDetail `json:"createPriceTable"`
}
// UpdatePriceTableParamDetail UpdatePriceTable detail param
type UpdatePriceTableParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePriceTableParam UpdatePriceTable request param
type UpdatePriceTableParam struct {
	BaseParam
	UpdatePriceTable UpdatePriceTableParamDetail `json:"updatePriceTable"`
}
// DeletePriceTableParamDetail DeletePriceTable detail param
type DeletePriceTableParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePriceTableParam DeletePriceTable request param
type DeletePriceTableParam struct {
	BaseParam
	DeletePriceTable DeletePriceTableParamDetail `json:"deletePriceTable"`
}
