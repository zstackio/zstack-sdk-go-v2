// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcVpnConnectionRemoteDetailParam DeleteVpcVpnConnectionRemote detail param
type DeleteVpcVpnConnectionRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnConnectionRemoteParam DeleteVpcVpnConnectionRemote request param
type DeleteVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params DeleteVpcVpnConnectionRemoteDetailParam `json:"params"`
}
