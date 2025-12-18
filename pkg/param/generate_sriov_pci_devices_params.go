// Copyright (c) ZStack.io, Inc.

package param

// GenerateSriovPciDevicesDetailParam GenerateSriovPciDevices详细参数
type GenerateSriovPciDevicesDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
	rest int `json:"virtPartNum" validate:"required"` // 必填
}

// GenerateSriovPciDevicesParam GenerateSriovPciDevices请求参数
type GenerateSriovPciDevicesParam struct {
	BaseParam
	Params GenerateSriovPciDevicesDetailParam `json:"params"` // 详细参数
}

