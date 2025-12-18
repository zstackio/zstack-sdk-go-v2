// Copyright (c) ZStack.io, Inc.

package param

// GCAliyunSnapshotRemoteDetailParam GCAliyunSnapshotRemote详细参数
type GCAliyunSnapshotRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// GCAliyunSnapshotRemoteParam GCAliyunSnapshotRemote请求参数
type GCAliyunSnapshotRemoteParam struct {
	BaseParam
	Params GCAliyunSnapshotRemoteDetailParam `json:"params"` // 详细参数
}

