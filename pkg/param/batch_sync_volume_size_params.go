// Copyright (c) ZStack.io, Inc.

package param

// BatchSyncVolumeSizeDetailParam BatchSyncVolumeSize detail param
type BatchSyncVolumeSizeDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// BatchSyncVolumeSizeParam BatchSyncVolumeSize request param
type BatchSyncVolumeSizeParam struct {
	BaseParam
	Params BatchSyncVolumeSizeDetailParam `json:"params"`
}
