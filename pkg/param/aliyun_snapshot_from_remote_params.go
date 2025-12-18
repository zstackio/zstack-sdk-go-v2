// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunSnapshotFromRemoteDetailParam DeleteAliyunSnapshotFromRemote详细参数
type DeleteAliyunSnapshotFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromRemoteParam DeleteAliyunSnapshotFromRemote请求参数
type DeleteAliyunSnapshotFromRemoteParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromRemoteDetailParam `json:"params"` // 详细参数
}

