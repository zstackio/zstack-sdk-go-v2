// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunSnapshotFromLocalDetailParam DeleteAliyunSnapshotFromLocal详细参数
type DeleteAliyunSnapshotFromLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromLocalParam DeleteAliyunSnapshotFromLocal请求参数
type DeleteAliyunSnapshotFromLocalParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromLocalDetailParam `json:"params"` // 详细参数
}

