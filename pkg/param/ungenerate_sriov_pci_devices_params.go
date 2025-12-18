// Copyright (c) ZStack.io, Inc.

package param

// UngenerateSriovPciDevicesDetailParam UngenerateSriovPciDevices detail param
type UngenerateSriovPciDevicesDetailParam struct {
	PciDeviceUuid string `json:"pciDeviceUuid" validate:"required"`
}

// UngenerateSriovPciDevicesParam UngenerateSriovPciDevices request param
type UngenerateSriovPciDevicesParam struct {
	BaseParam
	Params UngenerateSriovPciDevicesDetailParam `json:"params"`
}
