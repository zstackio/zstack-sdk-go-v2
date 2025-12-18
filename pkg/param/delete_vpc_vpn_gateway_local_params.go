// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcVpnGatewayLocalDetailParam DeleteVpcVpnGatewayLocal detail param
type DeleteVpcVpnGatewayLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnGatewayLocalParam DeleteVpcVpnGatewayLocal request param
type DeleteVpcVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcVpnGatewayLocalDetailParam `json:"params"`
}
