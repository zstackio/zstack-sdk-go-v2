// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualBorderRouterRemoteDetailParam UpdateVirtualBorderRouterRemote详细参数
type UpdateVirtualBorderRouterRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"localGatewayIp,omitempty"`
	rest string `json:"peerGatewayIp,omitempty"`
	rest string `json:"peeringSubnetMask,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vlanId,omitempty"`
	rest string `json:"circuitCode,omitempty"`
}

// UpdateVirtualBorderRouterRemoteParam UpdateVirtualBorderRouterRemote请求参数
type UpdateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params UpdateVirtualBorderRouterRemoteDetailParam `json:"params"` // 详细参数
}

