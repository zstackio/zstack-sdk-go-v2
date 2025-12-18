// Copyright (c) ZStack.io, Inc.

package param

// GCAliyunSnapshotRemoteDetailParam GCAliyunSnapshotRemote detail param
type GCAliyunSnapshotRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// GCAliyunSnapshotRemoteParam GCAliyunSnapshotRemote request param
type GCAliyunSnapshotRemoteParam struct {
	BaseParam
	Params GCAliyunSnapshotRemoteDetailParam `json:"params"`
}
