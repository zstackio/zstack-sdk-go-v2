// Copyright (c) ZStack.io, Inc.

package param

// DetachDataVolumeFromVmDetailParam DetachDataVolumeFromVm detail param
type DetachDataVolumeFromVmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmUuid string `json:"vmUuid,omitempty"`
}

// DetachDataVolumeFromVmParam DetachDataVolumeFromVm request param
type DetachDataVolumeFromVmParam struct {
	BaseParam
	Params DetachDataVolumeFromVmDetailParam `json:"params"`
}
