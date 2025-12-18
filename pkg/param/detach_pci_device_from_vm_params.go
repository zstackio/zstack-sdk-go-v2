// Copyright (c) ZStack.io, Inc.

package param

// DetachPciDeviceFromVmDetailParam DetachPciDeviceFromVm detail param
type DetachPciDeviceFromVmDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// DetachPciDeviceFromVmParam DetachPciDeviceFromVm request param
type DetachPciDeviceFromVmParam struct {
	BaseParam
	Params DetachPciDeviceFromVmDetailParam `json:"params"`
}
