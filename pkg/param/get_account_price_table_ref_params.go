// Copyright (c) ZStack.io, Inc.

package param

// GetAccountPriceTableRefDetailParam GetAccountPriceTableRef detail param
type GetAccountPriceTableRefDetailParam struct {
	TableUuid string `json:"tableUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// GetAccountPriceTableRefParam GetAccountPriceTableRef request param
type GetAccountPriceTableRefParam struct {
	BaseParam
	Params GetAccountPriceTableRefDetailParam `json:"params"`
}
