// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcUserVpnGatewayDetailParam UpdateVpcUserVpnGateway detail param
type UpdateVpcUserVpnGatewayDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcUserVpnGatewayParam UpdateVpcUserVpnGateway request param
type UpdateVpcUserVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcUserVpnGatewayDetailParam `json:"params"`
}
