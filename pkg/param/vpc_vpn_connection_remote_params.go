// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcVpnConnectionRemoteDetailParam CreateVpcVpnConnectionRemote详细参数
type CreateVpcVpnConnectionRemoteDetailParam struct {
	rest string `json:"userGatewayUuid" validate:"required"` // 必填
	rest string `json:"vpnGatewayUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"localCidr" validate:"required"` // 必填
	rest string `json:"remoteCidr" validate:"required"` // 必填
	rest bool `json:"active" validate:"required"` // 必填
	rest string `json:"ikeConfUuid" validate:"required"` // 必填
	rest string `json:"ipsecConfUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVpcVpnConnectionRemoteParam CreateVpcVpnConnectionRemote请求参数
type CreateVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params CreateVpcVpnConnectionRemoteDetailParam `json:"params"` // 详细参数
}

