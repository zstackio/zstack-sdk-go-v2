// Copyright (c) ZStack.io, Inc.

package param

// SyncHybridEipFromRemoteDetailParam SyncHybridEipFromRemote detail param
type SyncHybridEipFromRemoteDetailParam struct {
	Type string `json:"type" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncHybridEipFromRemoteParam SyncHybridEipFromRemote request param
type SyncHybridEipFromRemoteParam struct {
	BaseParam
	Params SyncHybridEipFromRemoteDetailParam `json:"params"`
}
