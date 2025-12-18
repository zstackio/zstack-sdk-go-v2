// Copyright (c) ZStack.io, Inc.

package param

// AttachPriceTableToAccountDetailParam AttachPriceTableToAccount detail param
type AttachPriceTableToAccountDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	TableUuid string `json:"tableUuid" validate:"required"`
}

// AttachPriceTableToAccountParam AttachPriceTableToAccount request param
type AttachPriceTableToAccountParam struct {
	BaseParam
	Params AttachPriceTableToAccountDetailParam `json:"params"`
}
