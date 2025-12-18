// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcUserVpnGatewayFromRemoteDetailParam SyncVpcUserVpnGatewayFromRemote detail param
type SyncVpcUserVpnGatewayFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcUserVpnGatewayFromRemoteParam SyncVpcUserVpnGatewayFromRemote request param
type SyncVpcUserVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcUserVpnGatewayFromRemoteDetailParam `json:"params"`
}
