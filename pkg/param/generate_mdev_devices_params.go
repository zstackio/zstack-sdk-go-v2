// Copyright (c) ZStack.io, Inc.

package param

// GenerateMdevDevicesDetailParam GenerateMdevDevices详细参数
type GenerateMdevDevicesDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
	rest string `json:"mdevSpecUuid" validate:"required"` // 必填
}

// GenerateMdevDevicesParam GenerateMdevDevices请求参数
type GenerateMdevDevicesParam struct {
	BaseParam
	Params GenerateMdevDevicesDetailParam `json:"params"` // 详细参数
}

