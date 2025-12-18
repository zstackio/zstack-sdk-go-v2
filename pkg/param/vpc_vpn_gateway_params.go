// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcVpnGatewayDetailParam UpdateVpcVpnGateway详细参数
type UpdateVpcVpnGatewayDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateVpcVpnGatewayParam UpdateVpcVpnGateway请求参数
type UpdateVpcVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcVpnGatewayDetailParam `json:"params"` // 详细参数
}

