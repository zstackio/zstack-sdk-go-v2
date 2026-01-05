// Copyright (c) ZStack.io, Inc.

package param

// GenerateHygonMdevDevicesDetailParam GenerateHygonMdevDevices detail param
type GenerateHygonMdevDevicesDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	MaxQemuNum int `json:"maxQemuNum" validate:"required"`
}

// GenerateHygonMdevDevicesParam GenerateHygonMdevDevices request param
type GenerateHygonMdevDevicesParam struct {
	BaseParam
	Params GenerateHygonMdevDevicesDetailParam `json:"params"`
}
