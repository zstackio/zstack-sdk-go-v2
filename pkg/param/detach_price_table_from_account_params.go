// Copyright (c) ZStack.io, Inc.

package param

// DetachPriceTableFromAccountDetailParam DetachPriceTableFromAccount detail param
type DetachPriceTableFromAccountDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	TableUuid string `json:"tableUuid" validate:"required"`
}

// DetachPriceTableFromAccountParam DetachPriceTableFromAccount request param
type DetachPriceTableFromAccountParam struct {
	BaseParam
	Params DetachPriceTableFromAccountDetailParam `json:"params"`
}
