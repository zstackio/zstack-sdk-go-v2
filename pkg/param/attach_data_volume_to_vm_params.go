// Copyright (c) ZStack.io, Inc.

package param

// AttachDataVolumeToVmDetailParam AttachDataVolumeToVm detail param
type AttachDataVolumeToVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// AttachDataVolumeToVmParam AttachDataVolumeToVm request param
type AttachDataVolumeToVmParam struct {
	BaseParam
	Params AttachDataVolumeToVmDetailParam `json:"params"`
}
