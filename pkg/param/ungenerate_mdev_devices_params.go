// Copyright (c) ZStack.io, Inc.

package param

// UngenerateMdevDevicesDetailParam UngenerateMdevDevices detail param
type UngenerateMdevDevicesDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
}

// UngenerateMdevDevicesParam UngenerateMdevDevices request param
type UngenerateMdevDevicesParam struct {
	BaseParam
	Params UngenerateMdevDevicesDetailParam `json:"params"`
}
