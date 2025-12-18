// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcUserVpnGatewayFromRemoteDetailParam SyncVpcUserVpnGatewayFromRemote详细参数
type SyncVpcUserVpnGatewayFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncVpcUserVpnGatewayFromRemoteParam SyncVpcUserVpnGatewayFromRemote请求参数
type SyncVpcUserVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcUserVpnGatewayFromRemoteDetailParam `json:"params"` // 详细参数
}

