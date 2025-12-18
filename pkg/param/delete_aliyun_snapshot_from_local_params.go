// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunSnapshotFromLocalDetailParam DeleteAliyunSnapshotFromLocal detail param
type DeleteAliyunSnapshotFromLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromLocalParam DeleteAliyunSnapshotFromLocal request param
type DeleteAliyunSnapshotFromLocalParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromLocalDetailParam `json:"params"`
}
