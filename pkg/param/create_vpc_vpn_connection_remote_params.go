// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcVpnConnectionRemoteDetailParam CreateVpcVpnConnectionRemote detail param
type CreateVpcVpnConnectionRemoteDetailParam struct {
	UserGatewayUuid string `json:"userGatewayUuid" validate:"required"`
	VpnGatewayUuid string `json:"vpnGatewayUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	LocalCidr string `json:"localCidr" validate:"required"`
	RemoteCidr string `json:"remoteCidr" validate:"required"`
	Active bool `json:"active" validate:"required"`
	IkeConfUuid string `json:"ikeConfUuid" validate:"required"`
	IpsecConfUuid string `json:"ipsecConfUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcVpnConnectionRemoteParam CreateVpcVpnConnectionRemote request param
type CreateVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params CreateVpcVpnConnectionRemoteDetailParam `json:"params"`
}
