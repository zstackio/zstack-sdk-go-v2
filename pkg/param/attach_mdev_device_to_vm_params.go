// Copyright (c) ZStack.io, Inc.

package param

// AttachMdevDeviceToVmDetailParam AttachMdevDeviceToVm detail param
type AttachMdevDeviceToVmDetailParam struct {
	MdevDeviceUuid string `json:"mdevDeviceUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// AttachMdevDeviceToVmParam AttachMdevDeviceToVm request param
type AttachMdevDeviceToVmParam struct {
	BaseParam
	Params AttachMdevDeviceToVmDetailParam `json:"params"`
}
