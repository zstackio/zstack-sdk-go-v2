// Copyright (c) ZStack.io, Inc.

package param

// UngenerateHygonMdevDevicesDetailParam UngenerateHygonMdevDevices detail param
type UngenerateHygonMdevDevicesDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// UngenerateHygonMdevDevicesParam UngenerateHygonMdevDevices request param
type UngenerateHygonMdevDevicesParam struct {
	BaseParam
	Params UngenerateHygonMdevDevicesDetailParam `json:"params"`
}
