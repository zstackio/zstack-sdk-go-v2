// Copyright (c) ZStack.io, Inc.

package param

// ValidateVolumeSnapshotChainDetailParam ValidateVolumeSnapshotChain detail param
type ValidateVolumeSnapshotChainDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ValidateVolumeSnapshotChainParam ValidateVolumeSnapshotChain request param
type ValidateVolumeSnapshotChainParam struct {
	BaseParam
	Params ValidateVolumeSnapshotChainDetailParam `json:"params"`
}
