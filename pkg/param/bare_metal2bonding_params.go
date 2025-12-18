// Copyright (c) ZStack.io, Inc.

package param

// CreateBareMetal2BondingDetailParam CreateBareMetal2Bonding详细参数
type CreateBareMetal2BondingDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest int `json:"mode" validate:"required"` // 必填
	rest string `json:"slaves" validate:"required"` // 必填
	rest string `json:"opts,omitempty"`
}

// CreateBareMetal2BondingParam CreateBareMetal2Bonding请求参数
type CreateBareMetal2BondingParam struct {
	BaseParam
	Params CreateBareMetal2BondingDetailParam `json:"params"` // 详细参数
}

