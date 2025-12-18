// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcVpnConnectionLocalDetailParam DeleteVpcVpnConnectionLocal detail param
type DeleteVpcVpnConnectionLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnConnectionLocalParam DeleteVpcVpnConnectionLocal request param
type DeleteVpcVpnConnectionLocalParam struct {
	BaseParam
	Params DeleteVpcVpnConnectionLocalDetailParam `json:"params"`
}
