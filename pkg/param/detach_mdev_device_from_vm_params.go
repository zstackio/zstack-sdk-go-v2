// Copyright (c) ZStack.io, Inc.

package param

// DetachMdevDeviceFromVmDetailParam DetachMdevDeviceFromVm detail param
type DetachMdevDeviceFromVmDetailParam struct {
	MdevDeviceUuid string `json:"mdevDeviceUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// DetachMdevDeviceFromVmParam DetachMdevDeviceFromVm request param
type DetachMdevDeviceFromVmParam struct {
	BaseParam
	Params DetachMdevDeviceFromVmDetailParam `json:"params"`
}
