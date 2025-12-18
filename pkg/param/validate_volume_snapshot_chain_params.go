// Copyright (c) ZStack.io, Inc.

package param

// ValidateVolumeSnapshotChainDetailParam ValidateVolumeSnapshotChain详细参数
type ValidateVolumeSnapshotChainDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ValidateVolumeSnapshotChainParam ValidateVolumeSnapshotChain请求参数
type ValidateVolumeSnapshotChainParam struct {
	BaseParam
	Params ValidateVolumeSnapshotChainDetailParam `json:"params"` // 详细参数
}

