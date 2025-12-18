// Copyright (c) ZStack.io, Inc.

package param

// GenerateMdevDevicesDetailParam GenerateMdevDevices detail param
type GenerateMdevDevicesDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
	MdevSpecUuid string `json:"mdevSpecUuid" validate:"required"`
}

// GenerateMdevDevicesParam GenerateMdevDevices request param
type GenerateMdevDevicesParam struct {
	BaseParam
	Params GenerateMdevDevicesDetailParam `json:"params"`
}
