// Copyright (c) ZStack.io, Inc.

package param

// GetDataVolumeAttachableVmDetailParam GetDataVolumeAttachableVm detail param
type GetDataVolumeAttachableVmDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// GetDataVolumeAttachableVmParam GetDataVolumeAttachableVm request param
type GetDataVolumeAttachableVmParam struct {
	BaseParam
	Params GetDataVolumeAttachableVmDetailParam `json:"params"`
}
