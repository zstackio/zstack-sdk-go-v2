// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcUserVpnGatewayLocalDetailParam DeleteVpcUserVpnGatewayLocal detail param
type DeleteVpcUserVpnGatewayLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayLocalParam DeleteVpcUserVpnGatewayLocal request param
type DeleteVpcUserVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayLocalDetailParam `json:"params"`
}
