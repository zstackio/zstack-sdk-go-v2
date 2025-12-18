// Copyright (c) ZStack.io, Inc.

package param

// UngenerateSeMdevDevicesDetailParam UngenerateSeMdevDevices detail param
type UngenerateSeMdevDevicesDetailParam struct {
	MttyDeviceUuid string `json:"mttyDeviceUuid" validate:"required"`
}

// UngenerateSeMdevDevicesParam UngenerateSeMdevDevices request param
type UngenerateSeMdevDevicesParam struct {
	BaseParam
	Params UngenerateSeMdevDevicesDetailParam `json:"params"`
}
