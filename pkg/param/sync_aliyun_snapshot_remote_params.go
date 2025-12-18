// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunSnapshotRemoteDetailParam SyncAliyunSnapshotRemote详细参数
type SyncAliyunSnapshotRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"snapshotId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncAliyunSnapshotRemoteParam SyncAliyunSnapshotRemote请求参数
type SyncAliyunSnapshotRemoteParam struct {
	BaseParam
	Params SyncAliyunSnapshotRemoteDetailParam `json:"params"` // 详细参数
}

