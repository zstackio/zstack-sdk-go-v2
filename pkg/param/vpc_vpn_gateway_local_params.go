// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcVpnGatewayLocalDetailParam DeleteVpcVpnGatewayLocal详细参数
type DeleteVpcVpnGatewayLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnGatewayLocalParam DeleteVpcVpnGatewayLocal请求参数
type DeleteVpcVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcVpnGatewayLocalDetailParam `json:"params"` // 详细参数
}

