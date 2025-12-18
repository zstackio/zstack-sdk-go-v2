// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcUserVpnGatewayLocalDetailParam DeleteVpcUserVpnGatewayLocal详细参数
type DeleteVpcUserVpnGatewayLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayLocalParam DeleteVpcUserVpnGatewayLocal请求参数
type DeleteVpcUserVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayLocalDetailParam `json:"params"` // 详细参数
}

