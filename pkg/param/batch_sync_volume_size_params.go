// Copyright (c) ZStack.io, Inc.

package param

// BatchSyncVolumeSizeDetailParam BatchSyncVolumeSize详细参数
type BatchSyncVolumeSizeDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// BatchSyncVolumeSizeParam BatchSyncVolumeSize请求参数
type BatchSyncVolumeSizeParam struct {
	BaseParam
	Params BatchSyncVolumeSizeDetailParam `json:"params"` // 详细参数
}

