// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunSnapshotDetailParam UpdateAliyunSnapshot详细参数
type UpdateAliyunSnapshotDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateAliyunSnapshotParam UpdateAliyunSnapshot请求参数
type UpdateAliyunSnapshotParam struct {
	BaseParam
	Params UpdateAliyunSnapshotDetailParam `json:"params"` // 详细参数
}

