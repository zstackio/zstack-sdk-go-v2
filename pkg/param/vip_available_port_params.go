// Copyright (c) ZStack.io, Inc.

package param

// GetVipAvailablePortDetailParam GetVipAvailablePort详细参数
type GetVipAvailablePortDetailParam struct {
	rest string `json:"vipUuid" validate:"required"` // 必填
	rest string `json:"protocolType" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVipAvailablePortParam GetVipAvailablePort请求参数
type GetVipAvailablePortParam struct {
	BaseParam
	Params GetVipAvailablePortDetailParam `json:"params"` // 详细参数
}

