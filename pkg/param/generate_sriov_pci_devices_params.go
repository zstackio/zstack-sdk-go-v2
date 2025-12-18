// Copyright (c) ZStack.io, Inc.

package param

// GenerateSriovPciDevicesDetailParam GenerateSriovPciDevices detail param
type GenerateSriovPciDevicesDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
	VirtPartNum int `json:"virtPartNum" validate:"required"`
}

// GenerateSriovPciDevicesParam GenerateSriovPciDevices request param
type GenerateSriovPciDevicesParam struct {
	BaseParam
	Params GenerateSriovPciDevicesDetailParam `json:"params"`
}
