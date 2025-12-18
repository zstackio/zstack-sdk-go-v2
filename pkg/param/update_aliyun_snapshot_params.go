// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunSnapshotDetailParam UpdateAliyunSnapshot detail param
type UpdateAliyunSnapshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunSnapshotParam UpdateAliyunSnapshot request param
type UpdateAliyunSnapshotParam struct {
	BaseParam
	Params UpdateAliyunSnapshotDetailParam `json:"params"`
}
