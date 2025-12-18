// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcUserVpnGatewayRemoteDetailParam DeleteVpcUserVpnGatewayRemote detail param
type DeleteVpcUserVpnGatewayRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayRemoteParam DeleteVpcUserVpnGatewayRemote request param
type DeleteVpcUserVpnGatewayRemoteParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayRemoteDetailParam `json:"params"`
}
