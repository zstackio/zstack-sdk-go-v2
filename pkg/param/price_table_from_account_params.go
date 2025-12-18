// Copyright (c) ZStack.io, Inc.

package param

// DetachPriceTableFromAccountDetailParam DetachPriceTableFromAccount详细参数
type DetachPriceTableFromAccountDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"tableUuid" validate:"required"` // 必填
}

// DetachPriceTableFromAccountParam DetachPriceTableFromAccount请求参数
type DetachPriceTableFromAccountParam struct {
	BaseParam
	Params DetachPriceTableFromAccountDetailParam `json:"params"` // 详细参数
}

