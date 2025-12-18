// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccountPriceTableBindingDetailParam ChangeAccountPriceTableBinding detail param
type ChangeAccountPriceTableBindingDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	TableUuid string `json:"tableUuid" validate:"required"`
}

// ChangeAccountPriceTableBindingParam ChangeAccountPriceTableBinding request param
type ChangeAccountPriceTableBindingParam struct {
	BaseParam
	Params ChangeAccountPriceTableBindingDetailParam `json:"params"`
}
