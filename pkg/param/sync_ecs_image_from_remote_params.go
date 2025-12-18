// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsImageFromRemoteDetailParam SyncEcsImageFromRemote detail param
type SyncEcsImageFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsImageFromRemoteParam SyncEcsImageFromRemote request param
type SyncEcsImageFromRemoteParam struct {
	BaseParam
	Params SyncEcsImageFromRemoteDetailParam `json:"params"`
}
