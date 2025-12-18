// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcUserVpnGatewayRemoteDetailParam DeleteVpcUserVpnGatewayRemote详细参数
type DeleteVpcUserVpnGatewayRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayRemoteParam DeleteVpcUserVpnGatewayRemote请求参数
type DeleteVpcUserVpnGatewayRemoteParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayRemoteDetailParam `json:"params"` // 详细参数
}

