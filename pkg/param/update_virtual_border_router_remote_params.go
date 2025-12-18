// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualBorderRouterRemoteDetailParam UpdateVirtualBorderRouterRemote detail param
type UpdateVirtualBorderRouterRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	LocalGatewayIp string `json:"localGatewayIp,omitempty"`
	PeerGatewayIp string `json:"peerGatewayIp,omitempty"`
	PeeringSubnetMask string `json:"peeringSubnetMask,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VlanId string `json:"vlanId,omitempty"`
	CircuitCode string `json:"circuitCode,omitempty"`
}

// UpdateVirtualBorderRouterRemoteParam UpdateVirtualBorderRouterRemote request param
type UpdateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params UpdateVirtualBorderRouterRemoteDetailParam `json:"params"`
}
