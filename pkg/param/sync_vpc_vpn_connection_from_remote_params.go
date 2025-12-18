// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcVpnConnectionFromRemoteDetailParam SyncVpcVpnConnectionFromRemote详细参数
type SyncVpcVpnConnectionFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnConnectionFromRemoteParam SyncVpcVpnConnectionFromRemote请求参数
type SyncVpcVpnConnectionFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnConnectionFromRemoteDetailParam `json:"params"` // 详细参数
}

