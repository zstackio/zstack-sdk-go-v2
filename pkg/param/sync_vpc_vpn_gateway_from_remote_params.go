// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcVpnGatewayFromRemoteDetailParam SyncVpcVpnGatewayFromRemote detail param
type SyncVpcVpnGatewayFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnGatewayFromRemoteParam SyncVpcVpnGatewayFromRemote request param
type SyncVpcVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnGatewayFromRemoteDetailParam `json:"params"`
}
