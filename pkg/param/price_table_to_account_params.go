// Copyright (c) ZStack.io, Inc.

package param

// AttachPriceTableToAccountDetailParam AttachPriceTableToAccount详细参数
type AttachPriceTableToAccountDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"tableUuid" validate:"required"` // 必填
}

// AttachPriceTableToAccountParam AttachPriceTableToAccount请求参数
type AttachPriceTableToAccountParam struct {
	BaseParam
	Params AttachPriceTableToAccountDetailParam `json:"params"` // 详细参数
}

