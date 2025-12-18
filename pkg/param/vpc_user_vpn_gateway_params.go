// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcUserVpnGatewayDetailParam UpdateVpcUserVpnGateway详细参数
type UpdateVpcUserVpnGatewayDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateVpcUserVpnGatewayParam UpdateVpcUserVpnGateway请求参数
type UpdateVpcUserVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcUserVpnGatewayDetailParam `json:"params"` // 详细参数
}

