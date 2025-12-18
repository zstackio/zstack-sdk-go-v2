// Copyright (c) ZStack.io, Inc.

package param

// AddPciDeviceSpecToVmInstanceDetailParam AddPciDeviceSpecToVmInstance detail param
type AddPciDeviceSpecToVmInstanceDetailParam struct {
	PciSpecUuid string `json:"pciSpecUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	PciDeviceNumber int `json:"pciDeviceNumber,omitempty"`
}

// AddPciDeviceSpecToVmInstanceParam AddPciDeviceSpecToVmInstance request param
type AddPciDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddPciDeviceSpecToVmInstanceDetailParam `json:"params"`
}
