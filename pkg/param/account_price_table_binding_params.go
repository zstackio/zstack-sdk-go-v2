// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccountPriceTableBindingDetailParam ChangeAccountPriceTableBinding详细参数
type ChangeAccountPriceTableBindingDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"tableUuid" validate:"required"` // 必填
}

// ChangeAccountPriceTableBindingParam ChangeAccountPriceTableBinding请求参数
type ChangeAccountPriceTableBindingParam struct {
	BaseParam
	Params ChangeAccountPriceTableBindingDetailParam `json:"params"` // 详细参数
}

