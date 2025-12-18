// Copyright (c) ZStack.io, Inc.

package param

// UpdatePriceTableDetailParam UpdatePriceTable详细参数
type UpdatePriceTableDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdatePriceTableParam UpdatePriceTable请求参数
type UpdatePriceTableParam struct {
	BaseParam
	Params UpdatePriceTableDetailParam `json:"params"` // 详细参数
}

