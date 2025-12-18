// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunSnapshotRemoteDetailParam SyncAliyunSnapshotRemote detail param
type SyncAliyunSnapshotRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	SnapshotId string `json:"snapshotId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunSnapshotRemoteParam SyncAliyunSnapshotRemote request param
type SyncAliyunSnapshotRemoteParam struct {
	BaseParam
	Params SyncAliyunSnapshotRemoteDetailParam `json:"params"`
}
