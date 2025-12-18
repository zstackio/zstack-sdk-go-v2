// Copyright (c) ZStack.io, Inc.

package param

// SyncDataCenterFromRemoteDetailParam SyncDataCenterFromRemote detail param
type SyncDataCenterFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncDataCenterFromRemoteParam SyncDataCenterFromRemote request param
type SyncDataCenterFromRemoteParam struct {
	BaseParam
	Params SyncDataCenterFromRemoteDetailParam `json:"params"`
}
