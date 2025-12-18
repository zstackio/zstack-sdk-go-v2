// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootVolumeDetailParam SetVmBootVolume detail param
type SetVmBootVolumeDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// SetVmBootVolumeParam SetVmBootVolume request param
type SetVmBootVolumeParam struct {
	BaseParam
	Params SetVmBootVolumeDetailParam `json:"params"`
}
