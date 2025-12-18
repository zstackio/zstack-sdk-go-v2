// Copyright (c) ZStack.io, Inc.

package param

// SyncVolumeSizeDetailParam SyncVolumeSize详细参数
type SyncVolumeSizeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncVolumeSizeParam SyncVolumeSize请求参数
type SyncVolumeSizeParam struct {
	BaseParam
	Params SyncVolumeSizeDetailParam `json:"params"` // 详细参数
}

