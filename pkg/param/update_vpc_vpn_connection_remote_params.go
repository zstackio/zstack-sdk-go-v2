// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcVpnConnectionRemoteDetailParam UpdateVpcVpnConnectionRemote detail param
type UpdateVpcVpnConnectionRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LocalCidr string `json:"localCidr,omitempty"`
	RemoteCidr string `json:"remoteCidr,omitempty"`
	Active bool `json:"active,omitempty"`
	IkeConfUuid string `json:"ikeConfUuid,omitempty"`
	IpsecConfUuid string `json:"ipsecConfUuid,omitempty"`
}

// UpdateVpcVpnConnectionRemoteParam UpdateVpcVpnConnectionRemote request param
type UpdateVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params UpdateVpcVpnConnectionRemoteDetailParam `json:"params"`
}
