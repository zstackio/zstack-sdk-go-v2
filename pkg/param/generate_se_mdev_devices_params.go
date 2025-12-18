// Copyright (c) ZStack.io, Inc.

package param

// GenerateSeMdevDevicesDetailParam GenerateSeMdevDevices detail param
type GenerateSeMdevDevicesDetailParam struct {
	MttyDeviceUuid string `json:"mttyDeviceUuid" validate:"required"`
	VirtPartNum int `json:"virtPartNum" validate:"required"`
}

// GenerateSeMdevDevicesParam GenerateSeMdevDevices request param
type GenerateSeMdevDevicesParam struct {
	BaseParam
	Params GenerateSeMdevDevicesDetailParam `json:"params"`
}
