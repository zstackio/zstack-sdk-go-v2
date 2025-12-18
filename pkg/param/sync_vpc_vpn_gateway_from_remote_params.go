// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcVpnGatewayFromRemoteDetailParam SyncVpcVpnGatewayFromRemote详细参数
type SyncVpcVpnGatewayFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnGatewayFromRemoteParam SyncVpcVpnGatewayFromRemote请求参数
type SyncVpcVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnGatewayFromRemoteDetailParam `json:"params"` // 详细参数
}

