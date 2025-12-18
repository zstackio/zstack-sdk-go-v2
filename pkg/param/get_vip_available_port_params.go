// Copyright (c) ZStack.io, Inc.

package param

// GetVipAvailablePortDetailParam GetVipAvailablePort detail param
type GetVipAvailablePortDetailParam struct {
	VipUuid string `json:"vipUuid" validate:"required"`
	ProtocolType string `json:"protocolType" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVipAvailablePortParam GetVipAvailablePort request param
type GetVipAvailablePortParam struct {
	BaseParam
	Params GetVipAvailablePortDetailParam `json:"params"`
}
