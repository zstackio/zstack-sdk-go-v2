// Copyright (c) ZStack.io, Inc.

package param

// GetVmAttachableDataVolumeDetailParam GetVmAttachableDataVolume detail param
type GetVmAttachableDataVolumeDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetVmAttachableDataVolumeParam GetVmAttachableDataVolume request param
type GetVmAttachableDataVolumeParam struct {
	BaseParam
	Params GetVmAttachableDataVolumeDetailParam `json:"params"`
}
