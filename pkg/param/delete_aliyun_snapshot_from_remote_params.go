// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunSnapshotFromRemoteDetailParam DeleteAliyunSnapshotFromRemote detail param
type DeleteAliyunSnapshotFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromRemoteParam DeleteAliyunSnapshotFromRemote request param
type DeleteAliyunSnapshotFromRemoteParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromRemoteDetailParam `json:"params"`
}
