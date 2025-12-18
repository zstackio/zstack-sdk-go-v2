// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcVpnGatewayDetailParam UpdateVpcVpnGateway detail param
type UpdateVpcVpnGatewayDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcVpnGatewayParam UpdateVpcVpnGateway request param
type UpdateVpcVpnGatewayParam struct {
	BaseParam
	Params UpdateVpcVpnGatewayDetailParam `json:"params"`
}
