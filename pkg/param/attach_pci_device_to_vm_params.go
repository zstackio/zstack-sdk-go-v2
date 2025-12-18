// Copyright (c) ZStack.io, Inc.

package param

// AttachPciDeviceToVmDetailParam AttachPciDeviceToVm detail param
type AttachPciDeviceToVmDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// AttachPciDeviceToVmParam AttachPciDeviceToVm request param
type AttachPciDeviceToVmParam struct {
	BaseParam
	Params AttachPciDeviceToVmDetailParam `json:"params"`
}
