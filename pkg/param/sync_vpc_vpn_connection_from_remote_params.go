// Copyright (c) ZStack.io, Inc.

package param

// SyncVpcVpnConnectionFromRemoteDetailParam SyncVpcVpnConnectionFromRemote detail param
type SyncVpcVpnConnectionFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnConnectionFromRemoteParam SyncVpcVpnConnectionFromRemote request param
type SyncVpcVpnConnectionFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnConnectionFromRemoteDetailParam `json:"params"`
}
