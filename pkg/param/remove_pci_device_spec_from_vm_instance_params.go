// Copyright (c) ZStack.io, Inc.

package param

// RemovePciDeviceSpecFromVmInstanceDetailParam RemovePciDeviceSpecFromVmInstance detail param
type RemovePciDeviceSpecFromVmInstanceDetailParam struct {
	PciSpecUuid string `json:"pciSpecUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// RemovePciDeviceSpecFromVmInstanceParam RemovePciDeviceSpecFromVmInstance request param
type RemovePciDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemovePciDeviceSpecFromVmInstanceDetailParam `json:"params"`
}
