// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcUserVpnGatewayRemoteDetailParam CreateVpcUserVpnGatewayRemote detail param
type CreateVpcUserVpnGatewayRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Ip string `json:"ip" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcUserVpnGatewayRemoteParam CreateVpcUserVpnGatewayRemote request param
type CreateVpcUserVpnGatewayRemoteParam struct {
	BaseParam
	Params CreateVpcUserVpnGatewayRemoteDetailParam `json:"params"`
}
